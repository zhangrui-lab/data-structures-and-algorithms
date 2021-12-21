package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplay(t *testing.T) {
	k1, k2, k3, k6, k8 := 1, 2, 3, 6, 8
	v1, v2, v3, v6, v8 := 10, 20, 30, 60, 80

	// 		k2
	//	   /
	//	  k1
	splay := NewSplay()
	splay.Insert(k1, v1)
	splay.Insert(k2, v2)
	assert.Equal(t, splay.Height(), 1, "splay.Height() != 1")
	assert.Equal(t, splay.root.value, v2, "splay.root.value != v2")

	// 		k1
	//		 \
	//		 k2
	v := splay.Search(k1)
	assert.Equal(t, v, v1, "splay.Search(k1) != v1")
	assert.Equal(t, splay.root.value, v1, "splay.root.value != v1")

	// 		k6
	//	    /
	//    k2
	//	  /
	//	 k1
	splay.Insert(k6, v6)
	assert.Equal(t, splay.root.value, v6, "splay.root.value != v6")
	assert.Equal(t, splay.Height(), 2, "splay.Height() != 2")
	e2 := splay.root.lc
	e1 := e2.lc
	assert.Equal(t, e1.value, v1, "e1.value != v1")
	assert.Equal(t, e2.value, v2, "e2.value != v2")

	//         k8
	//         /
	//       k6
	//       /
	//     k3
	//	   /
	//	 k2
	//	 /
	//	k1
	splay.Insert(k3, v3)
	splay.Insert(k8, v8)
	assert.Equal(t, splay.Height(), 4, "splay.Height() != 4")
	assert.Equal(t, splay.root.value, v8, "splay.root.value != v8")

	//			k1
	//		      \
	//		      k6
	//		    /    \
	//		   k2     k8
	//		    \
	//		    k3
	v = splay.Search(k1)
	assert.Equal(t, v, v1, "splay.Search(k1) != v1")
	assert.Equal(t, splay.Height(), 3, "splay.Height() != 3")
	e1 = splay.root
	e6 := e1.rc
	e2 = e6.lc
	e8 := e6.rc
	e3 := e2.rc
	assert.Nil(t, e1.lc, "e1.lc != nil")
	assert.Equal(t, e6.value, v6, "e6.value != v6")
	assert.Equal(t, e2.value, v2, "e2.value != v2")
	assert.Equal(t, e8.value, v8, "e8.value != v8")
	assert.Equal(t, e3.value, v3, "e3.value != v3")

	//			k1
	//		      \
	//		      k6
	//		    /    \
	//		   k2     k8
	//		    \
	//		    k3

	//			k1
	//		      \
	//		      k3
	//		    /    \
	//		   k2     k6
	//		          \
	//		           k8

	//			k3
	//		   /   \
	//		  k1    k6
	//		   \     \
	//		   k2     k8

	//			 k6
	//		   /   \
	//		  k1    k8
	//		   \
	//		   k2
	splay.Remove(k3)
	e6 = splay.root
	e1 = e6.lc
	e8 = e6.rc
	e2 = e1.rc
	assert.Equal(t, e6.value, v6, "e6.value != v6")
	assert.Equal(t, e1.value, v1, "e1.value != v1")
	assert.Equal(t, e8.value, v8, "e8.value != v8")
	assert.Equal(t, e2.value, v2, "e2.value != v2")
	assert.Equal(t, splay.Height(), 2, "splay.Height() != 2")

	splay.Clear()
	assert.Equal(t, splay.Empty(), true, "splay.Empty() != true")
}
