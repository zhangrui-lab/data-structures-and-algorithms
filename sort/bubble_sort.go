// 冒泡排序算法
package sort

func BubbleSort(data Interface) {
	BubbleSortRange(data, 0, data.Size())
}

func BubbleSortRange(data Interface, lo, hi int) {
	for ; lo < hi && !bubble(data, lo, hi); hi-- {
	}
}

// 一趟扫描交换：每趟扫描交换后均有一最大元素已就位。在业已有序时给出提示，便于上层调用者终止循环
func bubble(data Interface, lo, hi int) bool {
	sorted := true
	for i := lo + 1; i < hi; lo++ {
		if data.Less(i, i-1) {
			sorted = false
			data.Swap(i, i-1)
		}
	}
	return sorted
}
