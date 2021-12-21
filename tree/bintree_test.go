package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//			   	   25
//		  		 /    \
//		       18      47
//		      /  \    /  \
//		     9   21  30   62
//		    /   /     \   /
//		   5   19     40 51
//		    \
//           8
func TestBinTree(t *testing.T) {
	v5, v8, v9, v18, v19, v21, v25, v30, v40, v47, v51, v62 := 5, 8, 9, 18, 19, 21, 25, 30, 40, 47, 51, 62

	// insert
	tree := New()
	assert.Equal(t, tree.Size(), 0, "tree.Size() != 0")
	assert.Equal(t, tree.Empty(), true, "tree.Empty() != true")

	e25 := tree.InsertAsRoot(v25)
	assert.Equal(t, tree.Root(), e25, "tree.Root() != e25")
	assert.Equal(t, tree.Size(), 1, "tree.Size() != 1")
	assert.Equal(t, tree.Empty(), false, "tree.Empty() != false")

	e18 := tree.InsertAsLc(e25, v18)
	e47 := tree.InsertAsRc(e25, v47)
	e9 := tree.InsertAsLc(e18, v9)
	e21 := tree.InsertAsRc(e18, v21)
	e30 := tree.InsertAsLc(e47, v30)
	e62 := tree.InsertAsRc(e47, v62)
	e5 := tree.InsertAsLc(e9, v5)
	_ = tree.InsertAsLc(e21, v19)
	_ = tree.InsertAsRc(e30, v40)
	_ = tree.InsertAsLc(e62, v51)
	_ = tree.InsertAsRc(e5, v8)
	assert.Equal(t, tree.Size(), 12, "tree.Size() != 12")

	// travel
	level := "[25 18 47 9 21 30 62 5 19 40 51 8]"
	pre := "[25 18 9 5 8 21 19 47 30 40 62 51]"
	in := "[5 8 9 18 19 21 25 30 40 47 51 62]"
	post := "[8 5 9 19 21 18 40 30 51 62 47 25]"
	var str []string
	visit := func(key, value interface{}) {
		str = append(str, fmt.Sprintf("%v", value))
	}

	tree.TravelLevel(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), level, fmt.Sprintf("tree.travelLevel != %s", level))
	str = str[:0]

	tree.TravelPre(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), pre, fmt.Sprintf("tree.TravelPre != %s", pre))
	str = str[:0]

	tree.TravelIn(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), in, fmt.Sprintf("tree.TravelIn != %s", in))
	str = str[:0]

	tree.TravelPost(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), post, fmt.Sprintf("tree.TravelPost != %s", post))
	str = str[:0]

	tree18 := tree.Secede(e18)
	assert.Equal(t, tree18.Size(), 6, "tree18.Size() != 6")
	assert.Equal(t, tree.Size(), 6, "tree.Size() != 6")
	assert.Equal(t, tree.root.lc, (*BinNode)(nil), "tree.root.lc != nil")

	v := tree18.Remove(e5)
	assert.Equal(t, v, 2, "tree18.Remove(e5) != 2")
	tree18.TravelIn(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), "[9 18 19 21]", "tree18.TravelIn != [9 18 19 21]")
	str = str[:0]

	e18 = tree.AttachAsLC(tree18)
	assert.Equal(t, tree18.Size(), 0, "tree18.Size() != 0")
	assert.Equal(t, tree18.Empty(), true, "tree18.Empty() != true")

	assert.Equal(t, tree.Size(), 10, "tree.Size() != 10")

	tree.TravelIn(visit)
	assert.Equal(t, fmt.Sprintf("%v", str), "[9 18 19 21 25 30 40 47 51 62]", "tree18.TravelIn != [9 18 19 21 25 30 40 47 51 62]")
	str = str[:0]
}
