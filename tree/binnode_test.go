package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 构建一棵二叉树：
//			   	   25
//		  		 /    \
//		       18      47
//		      /  \    /  \
//		     9   21  30   62
//		    /   /     \   /
//		   5   19     40 51
//		    \
//           8
//
// 该二叉树的层序遍历为： 25 18 47 9 21 30 62 5 19 40 51 8
// 该二叉树的前序遍历为： 25 18 9 5 8 21 19 47 30 40 62 51
// 该二叉树的中序遍历为： 5 8 9 18 19 21 25 30 40 47 51 62
// 该二叉树的后序遍历为： 8 5 9 19 21 18 40 30 51 62 47 25
func createTree() *BinNode {
	v5, v8, v9, v18, v19, v21, v25, v30, v40, v47, v51, v62 := 5, 8, 9, 18, 19, 21, 25, 30, 40, 47, 51, 62

	e5 := newBinNode(v5, v5, nil, nil, nil)
	e8 := newBinNode(v8, v8, nil, nil, nil)
	e9 := newBinNode(v9, v9, nil, nil, nil)
	e18 := newBinNode(v18, v18, nil, nil, nil)
	e19 := newBinNode(v19, v19, nil, nil, nil)
	e21 := newBinNode(v21, v21, nil, nil, nil)
	e25 := newBinNode(v25, v25, nil, nil, nil)
	e30 := newBinNode(v30, v30, nil, nil, nil)
	e40 := newBinNode(v40, v40, nil, nil, nil)
	e47 := newBinNode(v47, v47, nil, nil, nil)
	e51 := newBinNode(v51, v51, nil, nil, nil)
	e62 := newBinNode(v62, v62, nil, nil, nil)

	e25.lc, e25.rc, e25.parent = e18, e47, nil
	e18.lc, e18.rc, e18.parent = e9, e21, e25
	e47.lc, e47.rc, e47.parent = e30, e62, e25
	e9.lc, e9.rc, e9.parent = e5, nil, e18
	e21.lc, e21.rc, e21.parent = e19, nil, e18
	e30.lc, e30.rc, e30.parent = nil, e40, e47
	e62.lc, e62.rc, e62.parent = e51, nil, e47
	e5.lc, e5.rc, e5.parent = nil, e8, e9
	e19.lc, e19.rc, e19.parent = nil, nil, e21
	e40.lc, e40.rc, e40.parent = nil, nil, e30
	e51.lc, e51.rc, e51.parent = nil, nil, e62
	e8.lc, e8.rc, e8.parent = nil, nil, e5

	return e25
}

func TestBasic(t *testing.T) {
	root := createTree()
	e47 := root.rc

	assert.Equal(t, root.isRoot(), true, "root.isRoot != true")
	e18 := root.lc
	assert.Equal(t, e18.isRoot(), false, "e18.isRoot != false")
	assert.Equal(t, e18.isLc(), true, "e18.isLc != true")
	assert.Equal(t, e18.isRc(), false, "e18.isRc != false")
	assert.Equal(t, e18.hasChild(), true, "e18.hasChild != true")
	assert.Equal(t, e18.hasBothChild(), true, "e18.hasBothChild != true")
	assert.Equal(t, e18.sibling(), e47, "e18.sibling != e47")

	e9 := e18.lc
	assert.Equal(t, e9.uncle(), e47, "e9.uncle != e47")

	e5 := e18.lc.lc
	assert.NotEqual(t, e5, nil, "e5 == nil")
	assert.Equal(t, e5.hasChild(), true, "e5.hasChild != true")
	assert.Equal(t, e5.hasBothChild(), false, "e5.hasBothChild != false")
	assert.Equal(t, e5.isLeaf(), false, "e5.isLeaf != false")

	e8 := e5.rc
	assert.Equal(t, e8.isLeaf(), true, "e8.isLeaf != true")
	assert.Equal(t, e8.isRc(), true, "e8.isRc != true")
	assert.Equal(t, e8.isRoot(), false, "e8.isRoot != false")
	assert.Equal(t, e8.sibling(), (*BinNode)(nil), "e8.sibling != nil")
	assert.Equal(t, e8.uncle(), (*BinNode)(nil), "e8.uncle != nil")

	e19 := e18.rc.lc
	e30 := e47.lc
	assert.Equal(t, e18.successor(), e19, "e18.succ != e19")
	assert.Equal(t, root.successor(), e30, "root.succ != e30")
	assert.Equal(t, e9.successor(), e18, "e9.succ != e18")
	assert.Equal(t, e8.successor(), e9, "e8.succ != e9")

	assert.Equal(t, e18.precursor(), e9, "e18.precursor() != e9")
	assert.Equal(t, e8.precursor(), e5, "e8.precursor() != e5")
	assert.Equal(t, e19.precursor(), e18, "e19.precursor() != e18")

	e40 := e47.lc.rc
	assert.NotEqual(t, e40, (*BinNode)(nil), "e40 == nil")
	assert.Equal(t, e40.successor(), e47, "e40.succ != e47")
}

func TestInsert(t *testing.T) {
	// 构建createTree中以18为根的子树
	v5, v8, v9, v18, v19, v21 := 5, 8, 9, 18, 19, 21
	e18 := newBinNode(v18, v18, nil, nil, nil)
	assert.Equal(t, e18.isRoot(), true, "e18.isRoot != true")
	assert.Equal(t, e18.size(), 1, "e18.size != 1")

	e9 := e18.insertLc(v9, v9)
	assert.Equal(t, e18.lc, e9, "e18.lc != e9")
	assert.Equal(t, e9.isLc(), true, "e9.isLc != true")
	assert.Equal(t, e18.size(), 2, "e18.size != 2")

	e21 := e18.insertRc(v21, v21)
	assert.Equal(t, e18.rc, e21, "e18.rc != e21")
	assert.Equal(t, e21.isRc(), true, "e21.isRc != true")
	assert.Equal(t, e18.size(), 3, "e18.size != 3")

	assert.Equal(t, e9.parent, e18, "e9.parent != e18")
	assert.Equal(t, e21.parent, e18, "e21.parent != e18")

	e5 := e9.insertLc(v5, v5)
	e8 := e5.insertRc(v8, v8)
	_ = e21.insertLc(v19, v19)
	assert.Equal(t, e18.size(), 6, "e18.size != 6")
	assert.Equal(t, e5.size(), 2, "e5.size != 2")
	assert.Equal(t, e5.isLeaf(), false, "e5.isLeaf != false")
	assert.Equal(t, e8.isLeaf(), true, "e8.isLeaf != true")
}

func TestTravel(t *testing.T) {
	level := "[25 18 47 9 21 30 62 5 19 40 51 8]"
	pre := "[25 18 9 5 8 21 19 47 30 40 62 51]"
	in := "[5 8 9 18 19 21 25 30 40 47 51 62]"
	post := "[8 5 9 19 21 18 40 30 51 62 47 25]"
	var str []string
	root := createTree()
	visit := func(key, value interface{}) {
		str = append(str, fmt.Sprintf("%v", value))
	}

	// 层序
	root.travelLevel(visit)
	assert.Equal(t, root.size(), 12, "tree size != 12")
	assert.Equal(t, fmt.Sprintf("%v", str), level, fmt.Sprintf("root.travelLevel != %s", level))
	str = str[:0]

	// 前序
	root.dfsPre(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), pre, fmt.Sprintf("root.travelPre != %s", pre))
	str = str[:0]

	root.stackPre1(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), pre, fmt.Sprintf("root.travelPre != %s", pre))
	str = str[:0]

	root.stackPre2(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), pre, fmt.Sprintf("root.travelPre != %s", pre))
	str = str[:0]

	root.morrisPre(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), pre, fmt.Sprintf("root.travelPre != %s", pre))
	str = str[:0]

	//// 中序
	root.dfsIn(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), in, fmt.Sprintf("root.travelIn != %s", in))
	str = str[:0]

	root.stackIn1(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), in, fmt.Sprintf("root.travelIn != %s", in))
	str = str[:0]

	root.stackIn2(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), in, fmt.Sprintf("root.travelIn != %s", in))
	str = str[:0]

	root.backtrackIn(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), in, fmt.Sprintf("root.travelIn != %s", in))
	str = str[:0]

	root.iterationIn(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), in, fmt.Sprintf("root.travelIn != %s", in))
	str = str[:0]

	root.morrisIn(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), in, fmt.Sprintf("root.travelIn != %s", in))
	str = str[:0]

	// 后序
	root.dfsPost(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), post, fmt.Sprintf("root.travelPost != %s", post))
	str = str[:0]

	root.stackPost(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), post, fmt.Sprintf("root.travelPost != %s", post))
	str = str[:0]

	root.morrisPost(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), post, fmt.Sprintf("root.travelPost != %s", post))
	str = str[:0]
}

func TestRotate(t *testing.T) {
	e25 := createTree()
	e18 := e25.lc
	e9 := e18.lc
	e21 := e18.rc
	e47 := e25.rc
	e30 := e47.lc
	e62 := e47.rc

	//					   47
	//					 /    \
	//			   	   25      62
	//		  		 /    \    /
	//		       18     30  51
	//		      /  \     \
	//		     9   21    40
	//		    /   /
	//		   5   19
	//		    \
	//           8
	e := e25.leftRotate()
	assert.Equal(t, e, e47, "e25.leftRotate() != e47")
	assert.Equal(t, e.isRoot(), true, "e.isRoot() != true")
	assert.Equal(t, e.lc, e25, "e.lc != e25")
	assert.Equal(t, e.rc, e62, "e.rc != e25")

	assert.Equal(t, e25.isLc(), true, "e25.isLc() != true")
	assert.Equal(t, e25.parent, e, "e25.parent != e")
	assert.Equal(t, e25.rc, e30, "e25.rc != e30")

	assert.Equal(t, e30.parent, e25, "e30.parent != e25")

	assert.Equal(t, e18.parent, e25, "e18.parent != e25")
	assert.Equal(t, e18.lc, e9, "e18.lc != e9")
	assert.Equal(t, e18.rc, e21, "e18.rc != e21")

	//					  47
	//					/    \
	//			   	  18      62
	//		  		/   \     /
	//		       9    25    51
	//		      /    /  \
	//		     5    21  30
	//		     \    /    \
	//		      8  19    40
	//
	e = e25.rightRotate()
	assert.Equal(t, e47.lc, e18, "e47.lc != e18")

	assert.Equal(t, e18.parent, e47, "e18.parent != e47")
	assert.Equal(t, e18.lc, e9, "e18.lc != e9")
	assert.Equal(t, e18.rc, e25, "e18.rc != e25")

	assert.Equal(t, e25.rc, e30, "e25.rc != e30")
	assert.Equal(t, e25.lc, e21, "e25.lc != e21")

	assert.Equal(t, e21.parent, e25, "e21.parent != e25")
}

func TestHeight(t *testing.T) {
	e25 := createTree()
	e18 := e25.lc
	e9 := e18.lc
	e5 := e9.lc
	assert.NotEqual(t, e5, nil, "e5 != nil")

	e5.updateHeight()
	assert.Equal(t, e9.getHeight(), 0, "e9.getHeight() != 0")
	assert.Equal(t, e5.getHeight(), 1, "e5.getHeight() != 1")
	assert.Equal(t, e5.lc.getHeight(), 0, "e5.lc.getHeight() != -1")

	e5.height = 0
	e5.updateHeightAbove()
	assert.Equal(t, e5.getHeight(), 1, "e5.getHeight() != 1")
	assert.Equal(t, e9.getHeight(), 2, "e9.getHeight() != 2")
	assert.Equal(t, e18.getHeight(), 3, "e18.getHeight() != 3")
	assert.Equal(t, e25.getHeight(), 4, "e25.getHeight() != 4")

}

func TestConnect34(t *testing.T) {
	e25 := createTree()
	e18 := e25.lc
	e9 := e18.lc
	_ = e18.rc
	e5 := e9.lc
	e := connect34(e5, e9, e18, e5.lc, e5.rc, e9.rc, e18.rc)
	assert.Equal(t, e, e9, "e != e9")
	assert.Equal(t, e9.lc, e5, "e9.lc != e5")
	assert.Equal(t, e9.rc, e18, "e9.ec != e18")
	assert.Equal(t, e5.parent, e9, "e5.parent != e9")
	assert.Equal(t, e18.parent, e9, "e18.parent != e9")
	assert.Equal(t, e18.lc, (*BinNode)(nil), "e18.lc != nil")

	e25 = createTree()
	e47 := e25.rc
	e30 := e47.lc
	e62 := e47.rc

	e = connect34(e25, e47, e62, e25.lc, e47.lc, e62.lc, e62.rc)
	assert.Equal(t, e, e47, "e != e47")
	assert.Equal(t, e47.lc, e25, "e47.lc != e25")
	assert.Equal(t, e47.rc, e62, "e47.rc != e62")
	assert.Equal(t, e25.rc, e30, "e25.rc != e30")
	assert.Equal(t, e30.parent, e25, "e30.parent != e25")

	e25 = createTree()
	e18 = e25.lc
	e21 := e18.rc
	e19 := e21.lc
	e = connect34(e18, e19, e21, e18.lc, e19.lc, e19.rc, e21.rc)
	assert.Equal(t, e, e19, "e != e19")
	assert.Equal(t, e19.lc, e18, "e19.lc != e18")
	assert.Equal(t, e19.rc, e21, "e19.rc != e21")
}
