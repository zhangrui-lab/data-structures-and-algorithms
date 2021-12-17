package sort

import (
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
	assert.Equal(t, vec.String(), "{4, 6, 7, 8, 8, 10}", "not eq {4, 6, 7, 8, 8, 10}")
}
