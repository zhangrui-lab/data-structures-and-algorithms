package huffman

import (
	"bytes"
	"data-structures-and-algorithms/bitmap"
	"data-structures-and-algorithms/priority_queue"
	"data-structures-and-algorithms/queue"
	"fmt"
)

var (
	nilChar = byte(0x00)
)

// huffNode 树节点
type huffNode struct {
	lc, rc *huffNode
	char   byte
	weight int
}

// BitInfo 字符编码位模式
type BitInfo struct {
	Len     int
	BitInfo *bitmap.BitMap
}

// Huffman Huffman编码树
type Huffman struct {
	root  *huffNode
	table map[byte]BitInfo
	data  []byte
}

// NewHuffman 根据输入数据构建Huffman编码树
func NewHuffman(data []byte) *Huffman {
	h := &Huffman{data: data}
	h.generate()
	return h
}

// Encode 数据编码
func (h *Huffman) Encode() BitInfo {
	sbi := BitInfo{Len: 0, BitInfo: bitmap.NewBitMap(0)}
	for i := 0; i < len(h.data); i++ {
		cbi := h.table[h.data[i]]
		for l, s := 0, cbi.Len; l < s; l++ {
			if cbi.BitInfo.Test(l) {
				sbi.BitInfo.Set(sbi.Len)
			} else {
				sbi.BitInfo.Clear(sbi.Len)
			}
			sbi.Len++
		}
	}
	return sbi
}

// Decode 解压长度为len的为数据（encode 通过初始化 data 产生）
func (h *Huffman) Decode(sbi BitInfo) string {
	node := h.root
	var buff bytes.Buffer
	for i := 0; i < sbi.Len; {
		if node.lc == nil && node.rc == nil {
			buff.WriteByte(node.char)
			node = h.root
			continue
		}
		if !sbi.BitInfo.Test(i) {
			node = node.lc
		} else {
			node = node.rc
		}
		i++
	}
	if node.lc == nil && node.rc == nil {
		buff.WriteByte(node.char)
	}
	return buff.String()
}

// 构建huffman树
func (h *Huffman) generate() {
	pq := priority_queue.NewCompBinHeap(func(a, b interface{}) int { // 权重越小者优先级越大
		return b.(*Huffman).root.weight - a.(*Huffman).root.weight
	})
	probability := make(map[byte]int)
	for i := 0; i < len(h.data); i++ {
		probability[h.data[i]]++
	}
	for k, v := range probability {
		pq.Insert(&Huffman{root: &huffNode{char: k, weight: v}})
	}
	for pq.Size() > 1 {
		t1 := pq.DelMax().(*Huffman)
		t2 := pq.DelMax().(*Huffman)
		pq.Insert(&Huffman{root: &huffNode{char: nilChar, weight: t1.root.weight + t2.root.weight, lc: t1.root, rc: t2.root}})
	}
	h.root = pq.DelMax().(*Huffman).root
	h.table = make(map[byte]BitInfo)
	h.generateTable(h.root, bitmap.NewBitMap(10), 0, h.table)
}

// 构建编码表
func (h *Huffman) generateTable(node *huffNode, bitmap *bitmap.BitMap, depth int, table map[byte]BitInfo) {
	if node == nil {
		return
	}
	if node.lc == nil && node.rc == nil {
		table[node.char] = BitInfo{Len: depth, BitInfo: bitmap.Clone(depth)}
	}
	if node.lc != nil {
		bitmap.Clear(depth)
		h.generateTable(node.lc, bitmap, depth+1, table)
	}
	if node.rc != nil {
		bitmap.Set(depth)
		h.generateTable(node.rc, bitmap, depth+1, table)
	}
}

// 打印huffman树
func (h *Huffman) print() {
	node := h.root
	if node == nil {
		return
	}
	que := queue.New()
	que.Push(node)
	fmt.Println("Huffman Tree:")
	for !que.Empty() {
		for i, size := 0, que.Size(); i < size; i++ {
			node = que.Pop().(*huffNode)
			if node.lc != nil {
				que.Push(node.lc)
			}
			if node.rc != nil {
				que.Push(node.rc)
			}
			fmt.Printf("%c:%d\t", node.char, node.weight)
		}
		fmt.Println()
	}
	fmt.Println("Huffman Table:")
	for k, bmp := range h.table {
		fmt.Printf("%c:%s\n", k, bmp.BitInfo.Output(bmp.Len))
	}
}
