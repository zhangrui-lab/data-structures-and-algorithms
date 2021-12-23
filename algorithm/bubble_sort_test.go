package algorithm

import (
	"fmt"
)

func ExampleBubbleSort() {

	nums := []interface{}{32, 12, 54, 23, 56, 0, 13, 8, 9}
	BubbleSortRange(nums, 0, 3)
	fmt.Println(nums)

	BubbleSortRange(nums, 5, 8, func(a, b interface{}) int { // 指定区间进行倒叙
		return b.(int) - a.(int)
	})
	fmt.Println(nums)

	BubbleSort(nums) // 切片整体排序
	fmt.Println(nums)

	// Output:
	// [12 32 54 23 56 0 13 8 9]
	// [12 32 54 23 56 13 8 0 9]
	// [0 8 9 12 13 23 32 54 56]
}
