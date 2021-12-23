package priority_queue

import "data-structures-and-algorithms/contract"

// heapNode 左氏堆树节点
type heapNode struct {
	item   interface{}
	lc, rc *heapNode
	npl    int //空节点路径长度
}

// 约定npl(x) = npl(null) = 0
func (n *heapNode) getNpl() int {
	if n == nil || n.lc == nil && n.rc == nil {
		return 0
	}
	return n.npl
}

// LeftHeap 基于二叉树，以左式堆形式实现的PQ
// 左式堆的整体结构呈单侧倾斜状；依照惯例，其中节点的分布均偏向左侧。也就是说，左式堆将不再如完全二叉堆那样满足结构性。 为各节点引入所谓的“空节点路
// 径长度”指标。节点x的空节点路径长度（null path length），记作npl(x)。若x为外部节点， 则约定npl(x) = npl(null) = 0。反之若x为内部节点，
// 则npl(x)可递归地定义为：npl(x) = 1 + min( npl(lc(x)), npl(rc(x)) ) 也就是说，节点x的npl值取决于其左、右孩子npl值中的小者。不难验证
// ：npl(x)既等于x到外部节点的最近距离，同时也等于以x为根的最大满子树的高度。 左式堆是处处满足“左倾性”的二叉堆，即任一内部节点x都满足
// npl(lc(x)) >= npl(rc(x))也就是说，就npl指标而言，任一内部节点的左孩子都不小于其右孩子。 也就是说，左式堆中每个节点的npl值，仅取决于其右孩
// 子。(“左孩子的npl值不小于右孩子”并不意味着“左孩子的高度必不小于右孩子”。)
// 从x出发沿右侧分支一直前行直至空节点，经过的通路称作其最右侧通路，记作rPath(x)。在左式堆中，尽管右孩子高度可能大于左孩子，但由“各节点npl值均决定
// 于其右孩子”这一事实不难发现，每个节点的npl值，应恰好等于其最右侧通路的长度。rPath(r)的终点必为全堆中深度最小的外部节点。
type LeftHeap struct {
	root       *heapNode
	size       int
	comparator contract.Comparator
}

// NewLeftHeap 新建空堆
func NewLeftHeap(cmps ...contract.Comparator) *LeftHeap {
	cmp := contract.DefaultComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	return &LeftHeap{comparator: cmp}
}

// NewLeftHeapFromSlice 从给定切片建堆
func NewLeftHeapFromSlice(data []interface{}, cmps ...contract.Comparator) *LeftHeap {
	cmp := contract.DefaultComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	heap := &LeftHeap{comparator: cmp}
	for _, v := range data {
		heap.Insert(v)
	}
	return heap
}

// Size 大小
func (h *LeftHeap) Size() int {
	return h.size
}

// Empty 判空
func (h *LeftHeap) Empty() bool {
	return h.Size() <= 0
}

// Merge 合并堆 o 并将其清空
func (h *LeftHeap) Merge(o *LeftHeap) *LeftHeap {
	h.root = h.merge(h.root, o.root)
	h.size += o.size
	o.root = nil
	o.size = 0
	return h
}

// Insert 按照比较器确定的优先级次序插入词条
func (h *LeftHeap) Insert(item interface{}) {
	h.size++
	node := &heapNode{item: item}
	h.root = h.merge(h.root, node)
}

// GetMax 取出优先级最高的词条
func (h *LeftHeap) GetMax() interface{} {
	if h.Empty() {
		return nil
	}
	return h.root.item
}

// DelMax 删除优先级最高的词条
func (h *LeftHeap) DelMax() interface{} {
	if h.Empty() {
		return nil
	}
	h.size--
	item := h.root.item
	h.root = h.merge(h.root.lc, h.root.rc)
	return item
}

// 执行堆节点合并
func (h *LeftHeap) merge(s, d *heapNode) *heapNode {
	if s == nil {
		return d
	}
	if d == nil {
		return s
	}
	if h.comparator(s.item, d.item) < 0 {
		s, d = d, s
	}
	s.rc = h.merge(s.rc, d)
	if s.lc.getNpl() < s.rc.getNpl() {
		s.lc, s.rc = s.rc, s.lc
	}
	s.npl = s.rc.getNpl()
	if s.rc != nil {
		s.npl++
	}
	return s
}
