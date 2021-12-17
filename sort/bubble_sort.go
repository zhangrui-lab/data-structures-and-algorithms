// Package sort 冒泡排序算法
package sort

import "data-structures-and-algorithms/types"

// BubbleSort 冒泡排序
func BubbleSort(data types.Interface) {
	BubbleSortRange(data, 0, data.Size())
}

// BubbleSortRange 对[lo, hi)区间执行冒泡排序
func BubbleSortRange(data types.Interface, lo, hi int) {
	for ; lo < hi && !bubble(data, lo, hi); hi-- {
	}
}

// 对 [lo, hi) 进行一趟扫描交换：每趟扫描交换后均有一最大元素已就位。在业已有序时给出提示，便于上层调用者终止循环
func bubble(data types.Interface, lo, hi int) bool {
	sorted := true
	for i := lo + 1; i < hi; i++ {
		if data.Less(i, i-1) {
			sorted = false
			data.Swap(i, i-1)
		}
	}
	return sorted
}
