package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := New() // empty stack
	assert.Equal(t, s.Size(), 0, "s.Size() != 0")
	assert.Equal(t, s.Empty(), true, "s.Empty() != true")
	assert.Equal(t, s.Top(), nil, "s.Top() != nil")
	assert.Equal(t, s.Pop(), nil, "s.Pop() != nil")

	v1, v2 := 1, 2
	s.Push(v1)
	s.Push(v2) // [1 2 >
	assert.Equal(t, s.Size(), 2, "s.Size() != 2")
	assert.Equal(t, s.Empty(), false, "s.Empty() != false")
	assert.Equal(t, s.Top(), v2, "s.Top() != v2")

	assert.Equal(t, s.Pop(), v2, "s.Pop() != v2") // [ 1 >
	assert.Equal(t, s.Top(), v1, "s.Top() != v1")

	assert.Equal(t, s.Pop(), v1, "s.Pop() != v1")         // [ >
	assert.Equal(t, s.Empty(), true, "s.Empty() != true") // [ >

	// 逆序输出
	nums := []float32{10.1, 20.2, 30.3, 40.5, 50.5, 60.6}
	for i := 0; i < len(nums); i++ {
		s.Push(nums[i])
	}
	assert.Equal(t, s.Size(), len(nums), "s.Size() != len(nums)")
	for i := 0; i < len(nums); i++ {
		assert.Equal(t, s.Pop(), nums[len(nums)-i-1], "s.Pop() != nums[len(nums)-i-1]")
	}
}
