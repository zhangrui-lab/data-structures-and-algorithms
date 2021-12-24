package huffman

import (
	"fmt"
)

func ExampleNewHuffman() {
	str := "gsdcvxzvdfgsfzxcdfgfsxczxcasacxz"
	h := NewHuffman([]byte(str))
	fmt.Println(h.Decode(h.Encode()))

	str = "flskfscnslfjldksljchvsdcbnxchoa;jdaxmkcbksdhdfadbkj23h3o1d89xck 1yckzcnidyodabdgasdjusxvcnsdfagda,mc asdgia"
	h = NewHuffman([]byte(str))
	fmt.Println(h.Decode(h.Encode()))

	str = "skfhlfdkd234了就VS能力u哦留校察看0让三年从还4了好久哦爱上多啊和"
	h = NewHuffman([]byte(str))
	fmt.Println(h.Decode(h.Encode()))

	// Output:
	// gsdcvxzvdfgsfzxcdfgfsxczxcasacxz
	// flskfscnslfjldksljchvsdcbnxchoa;jdaxmkcbksdhdfadbkj23h3o1d89xck 1yckzcnidyodabdgasdjusxvcnsdfagda,mc asdgia
	// skfhlfdkd234了就VS能力u哦留校察看0让三年从还4了好久哦爱上多啊和
}
