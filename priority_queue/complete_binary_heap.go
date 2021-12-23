package priority_queue

import "data-structures-and-algorithms/contract"

// CompBinHeap 完全二叉堆(完全二叉树的结构性 + 堆序性)
type CompBinHeap struct {
	items      []interface{}
	comparator contract.Comparator
}

// NewCompBinHeap 新建空堆
func NewCompBinHeap(cmps ...contract.Comparator) *CompBinHeap {
	cmp := contract.DefaultComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	return &CompBinHeap{comparator: cmp}
}

// FromSlice 从给定的切片中创建堆
func FromSlice(data []interface{}, cmps ...contract.Comparator) *CompBinHeap {
	cmp := contract.DefaultComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	heap := &CompBinHeap{comparator: cmp, items: make([]interface{}, len(data))}
	copy(heap.items, data)
	heap.heapify()
	return heap
}

// Size 大小
func (h *CompBinHeap) Size() int {
	return len(h.items)
}

// Empty 判空
func (h *CompBinHeap) Empty() bool {
	return h.Size() <= 0
}

// Insert 按照比较器确定的优先级次序插入词条
func (h *CompBinHeap) Insert(item interface{}) {
	h.items = append(h.items, nil)
	h.items[h.Size()-1] = item
	h.percolateUp(h.Size() - 1)
}

// GetMax 取出优先级最高的词条
func (h *CompBinHeap) GetMax() interface{} {
	if h.Empty() {
		return nil
	}
	return h.items[0]
}

// DelMax 删除优先级最高的词条
func (h *CompBinHeap) DelMax() interface{} {
	if h.Empty() {
		return nil
	}
	item := h.items[0]
	h.items[0], h.items[h.Size()-1] = h.items[h.Size()-1], h.items[0]
	h.items[h.Size()-1] = nil
	h.items = h.items[:h.Size()-1]
	h.percolateDown(0)
	return item
}

// percolateDown 对向量第r个元素实施下滤
func (h *CompBinHeap) percolateDown(r int) {
	for i := r; i < h.Size(); {
		j := h.properParent(i)
		if i == j {
			break
		}
		h.items[i], h.items[j] = h.items[j], h.items[i]
		i = j
	}
}

// percolateUp 对向量第r个元素实施上滤
func (h *CompBinHeap) percolateUp(r int) {
	for {
		p, b := h.parent(r)
		if !b || h.comparator(h.items[p], h.items[r]) >= 0 {
			break
		}
		h.items[p], h.items[r] = h.items[r], h.items[p]
		r = p
	}
}

// Floyd建堆算法
func (h *CompBinHeap) heapify() {
	for p := h.lastInternal(); p >= 0; p-- {
		h.percolateDown(p)
	}
}

// 节点i的左孩子
func (h *CompBinHeap) lc(i int) (int, bool) {
	lc := i<<1 + 1
	return lc, lc < h.Size()
}

// 节点i的右孩子
func (h *CompBinHeap) rc(i int) (int, bool) {
	rc := i<<1 + 2
	return rc, rc < h.Size()
}

// 节点i的父节点
func (h *CompBinHeap) parent(i int) (int, bool) {
	p := (i - 1) >> 1
	return p, p >= 0
}

// 是否为内部节点
func (h *CompBinHeap) lastInternal() int {
	p, _ := h.parent(h.Size() - 1)
	return p
}

// 父子（至多）三者中的大者
func (h *CompBinHeap) properParent(i int) (max int) {
	max = i
	j, b := h.lc(i)
	if !b {
		return
	}
	if h.comparator(h.items[max], h.items[j]) < 0 {
		max = j
	}
	j, b = h.rc(i)
	if b && h.comparator(h.items[max], h.items[j]) < 0 {
		max = j
	}
	return
}
