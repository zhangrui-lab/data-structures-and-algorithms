package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvl(t *testing.T) {
	k1, k2, k3, k4, k5, k6, k7, k8 := 1, 2, 3, 4, 5, 6, 7, 8
	v1, v2, v3, v4, v5, v6, v7, v8 := 10, 20, 30, 40, 50, 60, 70, 80
	avl := NewAvl()

	//  k1
	//	 \
	//    k2
	avl.Insert(k1, v1)
	assert.Equal(t, avl.Height(), 0, "avl.Height() != 0")
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
	assert.Equal(t, avl.String(), "{{1,10}, {2,20}, {3,30}, {4,40}, {5,50}}",
		"avl.String() != {{1,10}, {2,20}, {3,30}, {4,40}, {5,50}}")

	avl.Insert(k8, v8)
	//    	  k4
	//   	/   \
	//     k2    k5
	//    /  \    \
	//   k1  k3   k8
	assert.Equal(t, avl.Height(), 2, "avl.Height() != 2")
	assert.Equal(t, avl.String(), "{{1,10}, {2,20}, {3,30}, {4,40}, {5,50}, {8,80}}",
		"avl.String() != {{1,10}, {2,20}, {3,30}, {4,40}, {5,50}, {8,80}}")
	e4 := avl.root
	e2 := e4.lc
	e5 := e4.rc
	assert.Equal(t, e4.value, v4, "e4.value != v4")
	assert.Equal(t, e2.value, v2, "e2.value != v2")
	assert.Equal(t, e5.value, v5, "e5.value != v5")
	assert.Equal(t, e5.isLeaf(), false, "e5.isLeaf() != false")

	avl.Insert(k7, v7)
	//    	  k4
	//   	/    \
	//     k2    k7
	//    /  \   / \
	//   k1  k3 k5  k8
	assert.Equal(t, avl.Height(), 2, "avl.Height() != 2")
	assert.Equal(t, avl.String(), "{{1,10}, {2,20}, {3,30}, {4,40}, {5,50}, {7,70}, {8,80}}",
		"avl.String() != {{1,10}, {2,20}, {3,30}, {4,40}, {5,50}, {7,70}, {8,80}}")
	assert.Equal(t, e5.isLeaf(), true, "e5.isLeaf() != true")

	avl.Insert(k6, v6)
	//    	  k4
	//   	/    \
	//     k2    k7
	//    /  \   / \
	//   k1  k3 k5  k8
	//           \
	//           k6
	assert.Equal(t, avl.String(), "{{1,10}, {2,20}, {3,30}, {4,40}, {5,50}, {6,60}, {7,70}, {8,80}}",
		"avl.String() != {{1,10}, {2,20}, {3,30}, {4,40}, {5,50}, {6,60}, {7,70}, {8,80}}")

	v := avl.Search(k2)
	assert.Equal(t, v, v2, "avl.Search(k2) != v2")
	v = avl.Search(k7)
	assert.Equal(t, v, v7, "avl.Search(k7) != v7")
	v = avl.Search(k8)
	assert.Equal(t, v, v8, "avl.Search(v8) != k8")

	e7 := e5.parent
	assert.Equal(t, e7.value, v7, "e7.value != v7")

	// remove
	avl.Remove(k8)
	//    	  k4
	//   	/    \
	//     k2    k6
	//    /  \   / \
	//   k1  k3 k5  k7
	assert.NotEqual(t, e5.parent, e7, "e5.parent == e7")
	assert.Equal(t, e5.isLeaf(), true, "e5.isLeaf() != true")
	assert.Equal(t, avl.Height(), 2, "avl.Height() != 2")

	avl.Clear()
	assert.Equal(t, avl.Empty(), true, "avl.Empty() != true")
}
