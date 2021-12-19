package tree

import (
	"data-structures-and-algorithms/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBst(t *testing.T) {
	v1, v2, v3, v4, v5 := types.Int(1), types.Int(2), types.Int(3), types.Int(4), types.Int(5)
	bst := NewBst()
	assert.Equal(t, bst.Size(), 0, "empty_bst.Size() != 0")
	assert.Equal(t, bst.Empty(), true, "empty_bst.Empty() != true")

	// ----------------- insert -------------------------
	// v2
	e2 := bst.Insert(v2)
	assert.Equal(t, e2.Data, v2, "e2.Data != v2")
	assert.Equal(t, e2.isRoot(), true, "e2.isRoot() != true")

	e := bst.Search(v2)
	assert.Equal(t, e2, e, "bst.Search(v2) != e2")

	e = bst.Search(v1)
	assert.Equal(t, e, (*BinNode)(nil), "bst.Search(v1) != nil")

	//     v2
	//    /
	//  v1
	e1 := bst.Insert(v1)
	assert.Equal(t, e1.isLc(), true, "e1.isLc() != true")
	assert.Equal(t, e1.sibling(), (*BinNode)(nil), "e1.sibling() != nil")
	assert.Equal(t, e1, e2.lc, "e1 != e2.lc")
	assert.Equal(t, e1.isLeaf(), true, "e1.isLeaf() != true")

	//		 v2
	//		/ \
	//	   v1  v5
	e5 := bst.Insert(v5)
	assert.Equal(t, bst.Size(), 3, "bst.Size() != 3")
	assert.Equal(t, bst.Empty(), false, "bst.Empty() != false")
	assert.Equal(t, e5.sibling(), e1, "e5.sibling() != e1")
	assert.Equal(t, e5.isRc(), true, "e5.isRc() != true")

	assert.Equal(t, bst.String(), "{1, 2, 5}", "bst.String() != {1, 2, 5}")

	//		 	 v2
	//			/ \
	//	     v1    v5
	//	       	  /
	//          v3
	//           \
	//            v4
	e3 := bst.Insert(v3)
	e4 := bst.Insert(v4)
	assert.Equal(t, e3.isLc(), true, "e3.isLc() != true")
	assert.Equal(t, e3.rc, e4, "e3.rc != e4")
	assert.Equal(t, e4.isLeaf(), true, "e4.isLeaf() != true")
	assert.Equal(t, e3.isLeaf(), false, "e3.isLeaf() != false")
	assert.Equal(t, e4.isLeaf(), true, "e4.isLeaf() != true")
	assert.Equal(t, bst.String(), "{1, 2, 3, 4, 5}", "bst.String() != {1, 2, 3, 4, 5}")

	// ----------------- search -------------------------
	e = bst.Search(v4)
	assert.Equal(t, e, e4, "e != e4")
	assert.Equal(t, bst.hot, e3, "after bst.Search(v4), bst.hot != e3")

	e = bst.Search(v3)
	assert.Equal(t, e, e3, "e != e3")
	assert.Equal(t, bst.hot, e5, "after bst.Search(v3), bst.hot != e5")

	e = bst.Search(v2)
	assert.Equal(t, e, e2, "e != e2")
	assert.Equal(t, bst.hot, (*BinNode)(nil), "after bst.Search(v2), bst.hot != nil")

	// ----------------- remove -------------------------

	//		 	 v3
	//			/ \
	//	     v1    v5
	//	       	  /
	//           v4
	b := bst.Remove(v2)
	assert.Equal(t, b, true, "bst.Remove(v2) != true")
	b = bst.Remove(v2)
	assert.Equal(t, b, false, "double exec, bst.Remove(v2) != false")

	//assert.Equal(t, e2.lc, nil, "e2.lc != nil")
	//assert.Equal(t, e2.lc, e2.rc, "e2.lc != e2.rc")
	e = bst.Search(v3)
	assert.Equal(t, e.isRoot(), true, "e3.isRoot() != true")
	assert.Equal(t, e5.lc, e4, "e5.lc != e4")
	assert.Equal(t, bst.String(), "{1, 3, 4, 5}", "bst.String() != {1, 3, 4, 5}")
	assert.Equal(t, bst.Size(), 4, "bst.Size() != 4")

	//		v3
	//   	  \
	//         v4
	bst.Remove(v5)
	bst.Remove(v1)
	assert.Equal(t, bst.String(), "{3, 4}", "bst.String() != {3, 4}")
}
