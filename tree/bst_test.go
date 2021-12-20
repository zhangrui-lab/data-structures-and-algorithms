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
	bst.Insert(v2, v2)
	e2 := bst.root
	assert.Equal(t, e2.data.(Entry).key, v2, "e2.data.(Entry).key != v2")
	assert.Equal(t, e2.data.(Entry).value, v2, "e2.data.(Entry).value != v2")
	assert.Equal(t, e2.isRoot(), true, "e2.isRoot() != true")

	v, err := bst.Search(v2)
	assert.Equal(t, err, nil, "bst.Search(v2).err != nil")
	assert.Equal(t, v.(types.Int), v2, "bst.Search(v2).val != v2")

	v, err = bst.Search(v1)
	assert.NotEqual(t, err, nil, "bst.Search(v1).err == nil")

	//     v2
	//    /
	//  v1
	bst.Insert(v1, v1)
	e1 := e2.lc
	assert.Equal(t, e1.isLc(), true, "e1.isLc() != true")
	assert.Equal(t, e1.sibling(), (*BinNode)(nil), "e1.sibling() != nil")
	assert.Equal(t, e1, e2.lc, "e1 != e2.lc")
	assert.Equal(t, e1.isLeaf(), true, "e1.isLeaf() != true")

	//		 v2
	//		/ \
	//	   v1  v5
	bst.Insert(v5, v5)
	e5 := e2.rc
	assert.Equal(t, bst.Size(), 3, "bst.Size() != 3")
	assert.Equal(t, bst.Empty(), false, "bst.Empty() != false")
	assert.Equal(t, e5.sibling(), e1, "e5.sibling() != e1")
	assert.Equal(t, e5.isRc(), true, "e5.isRc() != true")

	assert.Equal(t, bst.String(), "{{1,1}, {2,2}, {5,5}}", "bst.String() != {{1,1}, {2,2}, {5,5}}")
	assert.Equal(t, bst.Height(), 1, "bst.Height() != 1")

	//		 	 v2
	//			/ \
	//	     v1    v5
	//	       	  /
	//          v3
	//           \
	//            v4
	bst.Insert(v3, v3)
	e3 := e5.lc
	bst.Insert(v4, v4)
	e4 := e3.rc
	assert.Equal(t, e3.isLc(), true, "e3.isLc() != true")
	assert.Equal(t, e3.rc, e4, "e3.rc != e4")
	assert.Equal(t, e4.isLeaf(), true, "e4.isLeaf() != true")
	assert.Equal(t, e3.isLeaf(), false, "e3.isLeaf() != false")
	assert.Equal(t, e4.isLeaf(), true, "e4.isLeaf() != true")
	assert.Equal(t, bst.String(), "{{1,1}, {2,2}, {3,3}, {4,4}, {5,5}}", "bst.String() != {{1,1}, {2,2}, {3,3}, {4,4}, {5,5}}")

	// ----------------- search -------------------------
	v, err = bst.Search(v4)
	assert.Equal(t, err, nil, "bst.Search(v4).err != nil")
	assert.Equal(t, v.(types.Int), v4, "bst.Search(v4).val != v4")

	v, err = bst.Search(v3)
	assert.Equal(t, err, nil, "bst.Search(v3).err != nil")
	assert.Equal(t, v.(types.Int), v3, "bst.Search(v3).val != v3")
	assert.Equal(t, bst.hot, e5, "after bst.Search(v3), bst.hot != e5")

	bst.Search(v2)
	assert.Equal(t, bst.hot, (*BinNode)(nil), "after bst.Search(v2), bst.hot != nil")

	assert.Equal(t, bst.Height(), 3, "bst.Height() != 3")

	// ----------------- remove -------------------------

	//		 	 v3
	//			/ \
	//	     v1    v5
	//	       	  /
	//           v4

	err = bst.Remove(v2)
	assert.Equal(t, err, nil, "bst.Remove(v2) != nil")
	err = bst.Remove(v2)
	assert.NotEqual(t, err, nil, "double exec, bst.Remove(v2) == nil")

	assert.Equal(t, bst.Height(), 2, "bst.Height() != 2")

	assert.Equal(t, e5.lc, e4, "e5.lc != e4")
	assert.Equal(t, bst.String(), "{{1,1}, {3,3}, {4,4}, {5,5}}", "bst.String() != {{1,1}, {3,3}, {4,4}, {5,5}}")
	assert.Equal(t, bst.Size(), 4, "bst.Size() != 4")

	//		v3
	//   	  \
	//         v4
	bst.Remove(v5)
	bst.Remove(v1)
	assert.Equal(t, bst.String(), "{{3,3}, {4,4}}", "bst.String() != {{3,3}, {4,4}}")

	assert.Equal(t, bst.Height(), 1, "bst.Height() != 1")

	bst.Clear()

	assert.Equal(t, bst.Size(), 0, "bst.Clear.Size() != 0")
	assert.Equal(t, bst.Empty(), true, "bst.Clear.Empty() != true")
	assert.Equal(t, bst.root, (*BinNode)(nil), "bst.Clear.root != nil")
	assert.Equal(t, bst.hot, (*BinNode)(nil), "bst.Clear.hot != nil")

}
