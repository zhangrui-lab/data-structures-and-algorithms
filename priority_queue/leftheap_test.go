package priority_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLeftHeap(t *testing.T) {
	h1k1, h1k2, h1k3, h1k4, h1k5 := 10, 20, 30, 40, 50
	h2k1, h2k2, h2k3, h2k4, h2k5 := 1, 8, 60, 70, 25
	h1 := NewLeftHeap()
	h2 := NewLeftHeap()

	h1.Insert(h1k3)
	h1.Insert(h1k2)
	h1.Insert(h1k1)
	h1.Insert(h1k5)
	h1.Insert(h1k4)

	h2.Insert(h2k4)
	h2.Insert(h2k3)
	h2.Insert(h2k1)
	h2.Insert(h2k5)
	h2.Insert(h2k2)

	assert.Equal(t, h1.Size(), 5, "h1.Size() != 5")
	assert.Equal(t, h2.Size(), 5, "h2.Size() != 5")

	assert.Equal(t, h1.GetMax(), 50, "h1.GetMax() != 50")
	assert.Equal(t, h2.GetMax(), 70, "h2.GetMax() != 70")

	assert.Equal(t, h1.DelMax(), 50, "h1.DelMax() != 50")
	assert.Equal(t, h2.DelMax(), 70, "h2.DelMax() != 70")

	assert.Equal(t, h1.DelMax(), 40, "h1.GetMax() != 40")
	assert.Equal(t, h2.DelMax(), 60, "h2.GetMax() != 60")

	assert.Equal(t, h1.Size(), 3, "h1.Size() != 4")
	assert.Equal(t, h2.Size(), 3, "h2.Size() != 4")

	h1 = h1.Merge(h2)
	assert.Equal(t, h1.Size(), 6, "h1.Size() != 8")
	assert.Equal(t, h2.Empty(), true, "h2.Empty() != true")

	assert.Equal(t, h1.DelMax(), 30, "h1.DelMax() != 30")
	assert.Equal(t, h1.DelMax(), 25, "h1.DelMax() != 25")
	assert.Equal(t, h1.DelMax(), 20, "h1.DelMax() != 20")
	assert.Equal(t, h1.DelMax(), 10, "h1.DelMax() != 10")
	assert.Equal(t, h1.DelMax(), 8, "h1.DelMax() != 8")
	assert.Equal(t, h1.DelMax(), 1, "h1.DelMax() != 1")

	assert.Equal(t, h1.Empty(), true, "h1.Empty() != true")

	nums := []interface{}{2, 30, 18}
	h1 = NewLeftHeapFromSlice(nums)
	assert.Equal(t, h1.DelMax(), 30, "h1.DelMax() != 30")
	assert.Equal(t, h1.DelMax(), 18, "h1.DelMax() != 18")
	assert.Equal(t, h1.DelMax(), 2, "h1.DelMax() != 2")
}
