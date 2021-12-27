package algorithm

import "data-structures-and-algorithms/contract"

// MergeSort 归并排序：分而治之
func MergeSort(data []interface{}, lo, hi int, comparators ...contract.Comparator) {
	if hi-lo < 2 { // 平凡情况
		return
	}
	mid := (lo + hi) >> 1
	MergeSort(data, lo, mid, comparators...)
	MergeSort(data, mid, hi, comparators...) // 分
	comparator := contract.DefaultComparator
	if len(comparators) > 0 {
		comparator = comparators[0]
	}
	merge(data, lo, mid, hi, comparator) // 治
}

// 合并算法
func merge(data []interface{}, lo, mid, hi int, cmp contract.Comparator) {
	for lo < mid && mid < hi {
		if cmp(data[lo], data[mid]) > 0 {
			e := data[mid]
			if mid < hi-1 {
				copy(data[mid:hi], data[mid+1:hi])
			}
			copy(data[lo+1:hi], data[lo:hi])
			data[lo] = e
			mid++
		}
		lo++
	}
}
