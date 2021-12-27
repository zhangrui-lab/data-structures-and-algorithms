package algorithm

import (
	"fmt"
)

func ExampleMergeSort() {
	nums := []interface{}{4, 7, 2, 8, 9, 34, 21, 74, 9, 0, 8, 1} // len:12

	MergeSort(nums, 0, 5)
	fmt.Println(nums)

	MergeSort(nums, 0, 5, func(a, b interface{}) int {
		return b.(int) - a.(int)
	})
	fmt.Println(nums)

	MergeSort(nums, 0, len(nums), func(a, b interface{}) int {
		return b.(int) - a.(int)
	})
	fmt.Println(nums)

	nums = nums[:0]
	MergeSort(nums, 0, 0)
	fmt.Println(nums)
	// Output:
	// [2 4 7 8 9 34 21 74 9 0 8 1]
	// [9 8 7 4 2 34 21 74 9 0 8 1]
	// [74 34 21 9 9 8 8 7 4 2 1 0]
	// []
}
