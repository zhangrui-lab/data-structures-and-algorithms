package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New() // empty queue
	assert.Equal(t, q.Size(), 0, "q.Size() != 0")
	assert.Equal(t, q.Empty(), true, "q.Empty() != true")
	assert.Equal(t, q.Front(), nil, "q.Front() != nil")
	assert.Equal(t, q.Back(), nil, "q.Back() != nil")

	v1, v2 := 1, 2
	q.Push(v1) // < 1 ]
	q.Push(v2) // < 1 2]
	assert.Equal(t, q.Size(), 2, "q.Size() != 2")
	assert.Equal(t, q.Empty(), false, "q.Empty() != false")
	assert.Equal(t, q.Front(), v1, "q.Front() != v1")
	assert.Equal(t, q.Back(), v2, "q.Back() != v2")

	assert.Equal(t, q.Pop(), v1, "q.Pop() != v1") // < 2 ]
	assert.Equal(t, q.Front(), v2, "q.Front() != v2")
	assert.Equal(t, q.Back(), v2, "q.Back() != v2")

	assert.Equal(t, q.Pop(), v2, "q.Pop() != v2")         // < ]
	assert.Equal(t, q.Empty(), true, "q.Empty() != true") // < ]

	// 逆序输出
	nums := []float32{10.1, 20.2, 30.3, 40.5, 50.5, 60.6}
	for i := 0; i < len(nums); i++ {
		q.Push(nums[i])
	}
	assert.Equal(t, q.Size(), len(nums), "q.Size() != len(nums)")
	for i := 0; i < len(nums); i++ {
		assert.Equal(t, q.Pop(), nums[i], "q.Pop() != nums[i]")
	}
}
