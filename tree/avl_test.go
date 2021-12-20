package tree

import (
	"data-structures-and-algorithms/types"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvl(t *testing.T) {
	k1, k2, k3, k4, k5 := types.Int(1), types.Int(2), types.Int(3), types.Int(4), types.Int(5)
	v1, v2, v3, v4, v5 := types.Int(10), types.Int(20), types.Int(30), types.Int(40), types.Int(50)
	avl := NewAvl()

	//  k1
	//	 \
	//    k2
	avl.Insert(k1, v1)
	avl.Insert(k2, v2)
	assert.Equal(t, avl.Height(), 1, "avl.Height() != 1")

	//    k2
	//   /  \
	//  k1   k4
	//      / \
	//     k3  k5
	avl.Insert(k3, v3)
	avl.Insert(k4, v4)
	avl.Insert(k5, v5)
	assert.Equal(t, avl.Height(), 2, "avl.Height() != 1")
	assert.Equal(t, avl.String(), "{{1,10}, {2,20}, {3,30}, {4,40}, {5,50}}", "avl.String() != {{1,10}, {2,20}, {3,30}, {4,40}, {5,50}}")

	fmt.Println(k3, v3, k4, v4, k5, v5)
}
