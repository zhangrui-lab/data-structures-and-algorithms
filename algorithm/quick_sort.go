package algorithm

import (
	"data-structures-and-algorithms/contract"
	"math/rand"
	"time"
)

// QuickSort 快速排序：与 MergeSort 算法的主要区别在于`分`与`治`的不同策略。快排侧重于分，归并侧重于治。
func QuickSort(data []interface{}, lo, hi int, comparators ...contract.Comparator) {
	if hi-lo < 2 {
		return
	}
	comparator := contract.DefaultComparator
	if len(comparators) > 0 {
		comparator = comparators[0]
	}
	pivot := partition(data, lo, hi, comparator) // 划分元素
	QuickSort(data, lo, pivot, comparators...)
	QuickSort(data, pivot+1, hi, comparators...)
}

// partition 划分算法 data[lo, pivot) <= data[pivot] <= data(pivot, hi)
func partition(data []interface{}, lo, hi int, cmp contract.Comparator) int {
	rand.Seed(time.Now().Unix())
	random := rand.Intn(hi-lo) + lo // 随机化算法：可选取更好的随机策略
	data[lo], data[random] = data[random], data[lo]
	i, elem := lo, data[lo]
	for j := lo + 1; j < hi; j++ {
		if cmp(data[j], elem) < 0 {
			e := data[j]
			if i < j-1 {
				copy(data[i+2:j+1], data[i+1:j])
			}
			data[i] = e
			i++
		}
	}
	data[i] = elem
	return i
}
