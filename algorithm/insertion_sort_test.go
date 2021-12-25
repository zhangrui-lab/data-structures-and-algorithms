package algorithm

import (
	"fmt"
)

func ExampleInsertionSort() {
	nums := []interface{}{4, 7, 2, 8, 9, 34, 21, 74, 9, 0, 8, 1} // len:12

	InsertionSort(nums, 0, 5)
	fmt.Println(nums)

	InsertionSort(nums, 0, 5, func(a, b interface{}) int {
		return b.(int) - a.(int)
	})
	fmt.Println(nums)

	InsertionSort(nums, 0, len(nums), func(a, b interface{}) int {
		return b.(int) - a.(int)
	})
	fmt.Println(nums)

	// Output:
	// [2 4 7 8 9 34 21 74 9 0 8 1]
	// [9 8 7 4 2 34 21 74 9 0 8 1]
	// [74 34 21 9 9 8 8 7 4 2 1 0]
}
