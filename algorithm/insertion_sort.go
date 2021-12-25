package algorithm

import "data-structures-and-algorithms/contract"

// InsertionSort 对 data[lo, hi) 进行插入排序
func InsertionSort(data []interface{}, lo, hi int, comparators ...contract.Comparator) {
	comparator := contract.DefaultComparator
	if len(comparators) > 0 {
		comparator = comparators[0]
	}
	for i := lo + 1; i < hi; i++ {
		findAndInsert(data, data[i], comparator, lo, i)
	}
}

// 在 data 的 [lo, hi) 中寻找适合 elem 插入的位置 i 。将原data[i, hi)的元素后移一位并将 elem 插入 i。
func findAndInsert(data []interface{}, elem interface{}, comparator contract.Comparator, lo, hi int) {
	index := hi - 1
	for ; index >= lo && comparator(data[index], elem) > 0; index-- { // data[lo-1]处为假想的最小值哨兵节点，故 index=lo-1时必定成功匹配。
		data[index+1] = data[index]
	}
	if index < lo || comparator(data[index], elem) <= 0 { // 当data[index] <= elem 时，在 index+1中插入
		index++
	}
	data[index] = elem
}
