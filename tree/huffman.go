// Package tree Huffman 编码树实现
package tree

import (
	"data-structures-and-algorithms/types"
	"fmt"
)

const innerhuffNode = byte(0x0)

// huffNode超字符
type huffNode struct {
	ch     byte    // 字符
	weight float64 // 权重
}

// Less 权重大小
func (c huffNode) Less(o types.Sortable) bool {
	switch o.(type) {
	case huffNode:
		return c.weight > o.(huffNode).weight
	default:
		panic(fmt.Sprintf("param o need type of *huffNode。 given: %v", o))
	}
}

// 新建超字符
func newHuffNode(ch byte, weight float64) huffNode {
	return huffNode{ch: ch, weight: weight}
}

// 输入字符统计
func inputStatistics(input []byte) map[byte]uint {
	nums := make(map[byte]uint, 1<<8)
	for x := range input {
		nums[byte(x)]++
	}
	return nums
}

// HuffmanTree 哈夫曼编码树
type HuffmanTree struct {
	root *BinNode
}

// NewHuffmanTree 构建HuffmanTree
func NewHuffmanTree(input []byte) *HuffmanTree {
	return &HuffmanTree{}
}

// merge 合并huffman树,并返回合并后的新树
func (t *HuffmanTree) merge(o *HuffmanTree) *HuffmanTree {
	x := newHuffNode(innerhuffNode, t.root.Data.(huffNode).weight+o.root.Data.(huffNode).weight)
	return &HuffmanTree{root: &BinNode{Data: x}}
}

func initForest() {
}
