package vector

import (
	"data-structures-and-algorithms/types"
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
	v1, v2, v3, v4 := types.Int(10), types.Int(5), types.Int(20), types.Int(40)
	vec := New(3)
	assert.Equal(t, vec.insert(0, v1), true, "insert != true")
	assert.Equal(t, vec.insert(vec.Size()+1, v2), false, "insert != fasle")
	assert.Equal(t, vec.Size(), 1, "size != 1")

	v, ok := vec.At(0)
	assert.Equal(t, v, v1, "vec[0] != 10")
	assert.Equal(t, ok, nil, "err != nil")

	_, ok = vec.At(1)
	assert.NotEqualf(t, ok, nil, "err == nil")

	v, ok = vec.Front()
	assert.Equal(t, v, v1, "vec[0] != 10")
	assert.Equal(t, ok, nil, "err != nil")

	v, ok = vec.Back()
	assert.Equal(t, v, v1, "vec[0] != 10")
	assert.Equal(t, ok, nil, "err != nil")

	vec.Push(v3)
	v, ok = vec.Pop()
	assert.Equal(t, v, v3, "vec[0] != 20")
	assert.Equal(t, ok, nil, "err != nil")

	vec.Push(v2)
	vec.Push(v3)

	vec.Assign(vec.Size()-1, v4)
	v, _ = vec.At(vec.Size() - 1)
	assert.Equal(t, v, v4, "vec.last != 40")

	vec.RemoveRange(0, 2)
	assert.Equal(t, vec.Size(), 1, "size != 1")

	vec.Clear()
	assert.Equal(t, vec.Size(), 0, "size != 0")
}

func TestVector_Find(t *testing.T) {
	v1, v2, v3, v4 := types.Int(10), types.Int(20), types.Int(30), types.Int(40)
	vec := CopySlice(v1, v2, v3, v4)
	assert.Equal(t, vec.Size(), 4, "size != 4")
	assert.Equal(t, vec.Capacity(), 4<<1, "cap != 6")

	r := vec.Find(v1)
	assert.Equal(t, r, 0, "r != 0")
	r = vec.Find(v3)
	assert.Equal(t, r, 2, "r != 2")
	r = vec.Find(types.Int(100))
	assert.Equal(t, r, -1, "r != -1")

	r = vec.Search(v2)
	assert.Equal(t, r, 1, "r != 1")
	r = vec.Search(types.Int(15))
	assert.Equal(t, r, 0, "r != 1")
	r = vec.Search(types.Int(7))
	assert.Equal(t, r, -1, "r != 0")
}

func TestVector_Deduplicate(t *testing.T) {
	v1, v2, v3, v4 := types.Int(10), types.Int(20), types.Int(30), types.Int(40)
	vec := CopySlice(v1, v1, v2, v2, v3, v3, v4)
	assert.Equal(t, vec.Size(), 7, "size != 6")
	assert.Equal(t, vec.Disordered(), 0, "Disordered != 0")
	assert.Equal(t, vec.Uniquify(), 3, "Uniquify != 3")

	vec = CopySlice(v1, v1, v2, v2, v3, v3, v4)
	vec.Scrambling()
	assert.NotEqual(t, vec.Disordered(), 0, "Disordered == 0")
	assert.NotEqual(t, vec.Deduplicate(), 3, "Deduplicate != 3")
}
