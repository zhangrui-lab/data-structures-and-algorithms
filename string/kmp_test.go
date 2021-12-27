package string

import (
	"fmt"
)

func ExampleKmpMatch() {
	text := "zhangsanzhangsansanzhang"
	pattern := "sanzh"
	index := KmpMatch(pattern, text)
	fmt.Println(index)
	if index != len(text) {
		fmt.Println(text[index : index+len(pattern)])
	}

	pattern = "sansa"
	index = KmpMatch(pattern, text)
	fmt.Println(index)
	if index != len(text) {
		fmt.Println(text[index : index+len(pattern)])
	}

	pattern = "skhsdkf"
	index = KmpMatch(pattern, text)
	fmt.Println(index)
	if index != len(text) {
		fmt.Println(text[index : index+len(pattern)])
	}

	// Output:
	// 5
	// sanzh
	// 13
	// sansa
	// 24
}
