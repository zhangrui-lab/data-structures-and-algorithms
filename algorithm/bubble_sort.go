// Package algorithm 冒泡排序算法
package algorithm

import (
	"data-structures-and-algorithms/contract"
)

// BubbleSort 冒泡排序
func BubbleSort(data []interface{}, comparators ...contract.Comparator) {
	BubbleSortRange(data, 0, len(data))
}

// BubbleSortRange 对[lo, hi)区间执行冒泡排序
func BubbleSortRange(data []interface{}, lo, hi int, comparators ...contract.Comparator) {
	comparator := contract.DefaultComparator
	if len(comparators) > 0 {
		comparator = comparators[0]
	}
	for ; lo < hi && !bubble(data, lo, hi, comparator); hi-- {
	}
}

// 对 [lo, hi) 进行一趟扫描交换：每趟扫描交换后均有一最大元素已就位。在业已有序时给出提示，便于上层调用者终止循环
func bubble(data []interface{}, lo, hi int, comparator contract.Comparator) bool {
	sorted := true
	for i := lo + 1; i < hi; i++ {
		if comparator(data[i], data[i-1]) < 0 {
			sorted = false
			data[i], data[i-1] = data[i-1], data[i]
		}
	}
	return sorted
}
