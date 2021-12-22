package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBtree(t *testing.T) {
	btree := NewBtree(4)
	assert.Equal(t, btree.Size(), 0, "btree.Size() != 0")
	assert.True(t, btree.Empty(), "btree.Empty() != true")

	var info string
	// 10
	btree.Insert(10)
	info = "[[10]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	// 10,20
	btree.Insert(20)
	info = "[[10 20]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	// 10,15,20
	btree.Insert(15)
	info = "[[10 15 20]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	     20
	//	    /  \
	//  10,15 	25
	btree.Insert(25)
	info = "[[20]\n[10 15] [25]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")
	//	     20
	//	    /  \
	// 5,10,15 	25
	btree.Insert(5)
	info = "[[20]\n[5 10 15] [25]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	      10,20
	//	    /   |    \
	//   0,5   15    25
	btree.Insert(0)
	info = "[[10 20]\n[0 5] [15] [25]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	      10,20
	//	    /   |   \
	//   0,5   15   25,30
	btree.Insert(30)
	info = "[[10 20]\n[0 5] [15] [25 30]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	      10,20
	//	    /   |   \
	//   0,5   15   25,30, 40
	btree.Insert(40)
	info = "[[10 20]\n[0 5] [15] [25 30 40]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	      10,20
	//	    /   |   \
	//   0,5  12,15   25,30, 40
	btree.Insert(12)
	info = "[[10 20]\n[0 5] [12 15] [25 30 40]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	      10,20,40
	//	    /   |  \    \
	//   0,5  12,15 25,30 56
	btree.Insert(56)
	info = "[[10 20 40]\n[0 5] [12 15] [25 30] [56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	       10,20,40
	//	     /   |  \    \
	// 0,5,8  12,15 25,30 56
	btree.Insert(8)
	info = "[[10 20 40]\n[0 5 8] [12 15] [25 30] [56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	       10,20,40
	//	     /   |   \     \
	// 0,5,8  11,12,15 25,30 56
	btree.Insert(11)
	info = "[[10 20 40]\n[0 5 8] [11 12 15] [25 30] [56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	       10,20,40
	//	     /   |   \     \
	//  5,8  11,12,15 25,30 56
	btree.Remove(0)
	info = "[[10 20 40]\n[5 8] [11 12 15] [25 30] [56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	       10,20,40
	//	     /   |   \     \
	//   8  11,12,15 25,30 56
	btree.Remove(5)
	info = "[[10 20 40]\n[8] [11 12 15] [25 30] [56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	       11,20,40
	//	     /   |   \     \
	//     10  12,15 25,30 56
	btree.Remove(8)
	info = "[[11 20 40]\n[10] [12 15] [25 30] [56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	assert.Equal(t, btree.Height(), 1, "btree.Height() != 1")
	assert.Equal(t, btree.Size(), 9, "btree.Size() != 9")
	assert.Equal(t, btree.Order(), 4, "btree.Order() != 4")

	//	       12,20,40
	//	     /   |   \  \
	//     11  15  25,30 56
	btree.Remove(10)
	info = "[[12 20 40]\n[11] [15] [25 30] [56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	       12,25,40
	//	     /   |  \  \
	//      11  15   30 56
	btree.Remove(20)
	info = "[[12 25 40]\n[11] [15] [30] [56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	assert.Equal(t, btree.Search(15).(int), 15, "btree.Search(15) != 15")
	assert.Equal(t, btree.Search(11).(int), 11, "btree.Search(11) != 11")
	assert.Equal(t, btree.Height(), 1, "btree.Height() != 1")
	assert.Equal(t, btree.Size(), 7, "btree.Size() != 7")

	//	       12,25
	//	     /   |  \
	//      11  15   30,56
	btree.Remove(40)
	info = "[[12 25]\n[11] [15] [30 56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	        25
	//	     /     \
	//     12,15   30,56
	btree.Remove(11)
	info = "[[25]\n[12 15] [30 56]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	        25
	//	     /     \
	//     12,15   30
	btree.Remove(56)
	info = "[[25]\n[12 15] [30]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	//	        25
	//	     /     \
	//     15       30
	btree.Remove(12)
	info = "[[25]\n[15] [30]]"
	assert.Equal(t, btree.levelInfo(), info, "btree.levelInfo() != info")

	btree.Clear()
	assert.Equal(t, btree.Size(), 0, "btree.Size() != 0")
	assert.True(t, btree.Empty(), "btree.Empty() != true")
}
