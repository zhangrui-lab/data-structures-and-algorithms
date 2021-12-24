package algorithm

import "fmt"

type user struct {
	id   int
	age  int
	name string
}

func ExampleSelectionSort() {
	nums := []interface{}{32, 12, 54, 23, 56, 0, 13, 8, 9}
	SelectionSort(nums, 0, 3)
	fmt.Println(nums)

	SelectionSort(nums, 5, 8, func(a, b interface{}) int { // 指定区间进行倒叙
		return b.(int) - a.(int)
	})
	fmt.Println(nums)

	users := []interface{}{
		user{id: 1, age: 23, name: "李四1"},
		user{id: 7, age: 43, name: "李四7"},
		user{id: 2, age: 25, name: "李四2"},
		user{id: 5, age: 23, name: "李四5"},
		user{id: 6, age: 73, name: "李四6"},
		user{id: 3, age: 21, name: "李四3"},
		user{id: 4, age: 23, name: "李四4"},
	}
	fmt.Println(users)
	SelectionSort(users, 0, 3, func(a, b interface{}) int {
		return a.(user).age - b.(user).age
	})
	fmt.Println(users)

	SelectionSort(users, 3, 7, func(a, b interface{}) int {
		return a.(user).id - b.(user).id
	})
	fmt.Println(users)

	// Output:
	// [12 32 54 23 56 0 13 8 9]
	// [12 32 54 23 56 13 8 0 9]
	//[{1 23 李四1} {7 43 李四7} {2 25 李四2} {5 23 李四5} {6 73 李四6} {3 21 李四3} {4 23 李四4}]
	//[{1 23 李四1} {2 25 李四2} {7 43 李四7} {5 23 李四5} {6 73 李四6} {3 21 李四3} {4 23 李四4}]
	//[{1 23 李四1} {2 25 李四2} {7 43 李四7} {3 21 李四3} {4 23 李四4} {5 23 李四5} {6 73 李四6}]

}
