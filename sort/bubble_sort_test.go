package sort

import (
	"data-structures-and-algorithms/list"
	"data-structures-and-algorithms/types"
	"data-structures-and-algorithms/vector"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort_vector(t *testing.T) {
	v1, v2, v3, v4, v5, v6 := types.Int(4), types.Int(7), types.Int(6), types.Int(8), types.Int(8), types.Int(10)
	vec := vector.CopySlice(v1, v2, v3, v4, v5, v6)
	vec.Scrambling()
	assert.NotEqual(t, vec.Disordered(), 0, "Disordered == 0")
	BubbleSort(vec)
	assert.Equal(t, vec.Disordered(), 0, "Disordered != 0")
	assert.Equal(t, vec.String(), "{4, 6, 7, 8, 8, 10}", "vec.String() != {4, 6, 7, 8, 8, 10}")
}

func TestBubbleSort_list(t *testing.T) {
	v1, v2, v3, v4, v5 := types.Int(1), types.Int(2), types.Int(3), types.Int(4), types.Int(5)
	l := list.New()
	l.PushBack(v2)
	l.PushBack(v3)
	l.PushBack(v1)
	l.PushBack(v5)
	l.PushBack(v4) // {2, 3, 1, 5, 4}
	assert.Equal(t, l.String(), "{2, 3, 1, 5, 4}", "l.String() != {2, 3, 1, 5, 4}")

	BubbleSort(l) // {1, 2, 3, 4, 5}
	assert.Equal(t, l.String(), "{1, 2, 3, 4, 5}", "l.String() != {1, 2, 3, 4, 5}")
}
