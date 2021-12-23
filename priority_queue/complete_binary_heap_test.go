package priority_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompBinHeap(t *testing.T) {
	nums := []interface{}{3, 45, 12, 5, 7, 2, 56, 31, 14, 67}
	pq := FromSlice(nums)
	assert.Equal(t, pq.Size(), 10, "pq.Size() != 10")
	assert.Equal(t, pq.Empty(), false, "pq.Empty() != false")
	assert.Equal(t, pq.GetMax(), 67, "pq.GetMax() != 67")

	assert.Equal(t, pq.DelMax(), 67, "pq.DelMax() != 67")
	assert.Equal(t, pq.DelMax(), 56, "pq.DelMax() != 56")
	assert.Equal(t, pq.DelMax(), 45, "pq.DelMax() != 45")

	pq.Insert(100)
	assert.Equal(t, pq.DelMax(), 100, "pq.DelMax() != 100")
	assert.Equal(t, pq.Size(), 7, "pq.Size() != 7")

	pq.Insert(88)
	pq.Insert(99)
	assert.Equal(t, pq.DelMax(), 99, "pq.DelMax() != 99")
	assert.Equal(t, pq.DelMax(), 88, "pq.DelMax() != 88")
	assert.Equal(t, pq.DelMax(), 31, "pq.DelMax() != 31")
}
