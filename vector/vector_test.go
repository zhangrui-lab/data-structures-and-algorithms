package vector

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVector_New(t *testing.T) {
	vec := New(10)
	assert.Equal(t, vec.Size(), 0, "vec.New(10) size except 0")
	assert.Equal(t, vec.Capacity(), 10, "vec.New(10) cap except 10")
	assert.Equal(t, vec.Empty(), true, "vec.New(10).Empty() except true")
}

func TestVector_Insert(t *testing.T) {
	v1, v2 := 10, 5
	vec := New(3)
	assert.Equal(t, vec.insert(0, v1), true, "insert != true")
	assert.Equal(t, vec.insert(vec.Size()+1, v2), false, "insert != fasle")
	assert.Equal(t, vec.Size(), 1, "size != 1")

	v := vec.At(0)
	assert.Equal(t, v, v1, "vec[0] != 10")

	v = vec.Front()
	assert.Equal(t, v, v1, "vec[0] != 10")

	v = vec.Back()
	assert.Equal(t, v, v1, "vec[0] != 10")
}

func TestVector_Copy_Clear(t *testing.T) {
	v1, v2, v3, v4 := 10, 20, 30, 40
	vec1 := FromSlice(v1, v2, v3, v4)
	vec2 := Copy(vec1)
	assert.Equal(t, vec2.Size(), 4, "size != 4")
	assert.Equal(t, vec2.Capacity(), 4, "cap != 4")
	vec2.Clear()
	assert.Equal(t, vec2.Size(), 0, "size != 4")
	assert.Equal(t, vec2.Capacity(), 4, "cap != 4")
}

func TestVector_Remove(t *testing.T) {
	v1, v2, v3, v4 := 10, 20, 30, 40
	vec := FromSlice(v1, v2, v3, v4)

	e := vec.Remove(0)
	assert.Equal(t, e, v1, "e != 10")
}

func TestVector_Deduplicate(t *testing.T) {
	v1, v2, v3, v4 := 10, 20, 30, 40
	vec := FromSlice(v1, v1, v2, v2, v3, v3, v4)
	assert.Equal(t, vec.Size(), 7, "size != 6")
	assert.Equal(t, vec.Uniquify(), 3, "Uniquify != 3")

	vec = FromSlice(v1, v1, v2, v2, v3, v3, v4)
	vec.Scrambling()
	assert.NotEqual(t, vec.Deduplicate(), 3, "Deduplicate != 3")
}

func TestIterator(t *testing.T) {
	nums := []interface{}{10, 20, 30, 40, 50, 60}
	vec := FromSlice(nums...)
	i := 0
	for iter := vec.Begin(); !iter.Equal(vec.End()); iter.Next() {
		assert.Equal(t, iter.Valid(), true, "iter.Valid() != true")
		assert.Equal(t, iter.Value().(int), nums[i])
		i++
	}

	iter := vec.Begin().Forward(3)
	assert.Equal(t, iter.Valid(), true, "iter.Valid() != true")
	assert.Equal(t, iter.Value().(int), nums[3])
}
