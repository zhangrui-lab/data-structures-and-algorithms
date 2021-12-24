package algorithm

import "data-structures-and-algorithms/contract"

// SelectionSort 选择排序
func SelectionSort(data []interface{}, lo, hi int, comparators ...contract.Comparator) {
	comparator := contract.DefaultComparator
	if len(comparators) > 0 {
		comparator = comparators[0]
	}
	for lo < hi {
		i := doSelect(data, lo, hi, comparator)
		data[hi-1], data[i] = data[i], data[hi-1]
		hi--
	}
}

// 区间 [lo, hi) 最大值选择
// 将选择算法替换为 O(lgn) 的堆极值获取时,即可转化为堆排序
func doSelect(data []interface{}, lo, hi int, comparator contract.Comparator) int {
	max := lo
	for lo = lo + 1; lo < hi; lo++ {
		if comparator(data[max], data[lo]) <= 0 { // 优先考虑靠后者
			max = lo
		}
	}
	return max
}
