package algorithm

import (
	"data-structures-and-algorithms/contract"
	"data-structures-and-algorithms/priority_queue"
)

// HeapSort 对 data[lo, hi) 进行排序
func HeapSort(data []interface{}, lo, hi int, comparators ...contract.Comparator) {
	comparator := contract.DefaultComparator
	if len(comparators) > 0 {
		comparator = comparators[0]
	}
	pq := priority_queue.NewCompBinHeapFromSlice(data[lo:hi], comparator) // 使用自定义heap结构，丧失了就地排序的特性
	for i := hi - 1; i >= lo; i-- {
		data[i] = pq.DelMax()
	}
}
