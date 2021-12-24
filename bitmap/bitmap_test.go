package bitmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitMap(t *testing.T) {
	size := 16
	bm := NewBitMap(size)
	assert.Equal(t, bm.Output(size), "0000000000000000", "bm.Output(size) != 0000000000000000")

	assert.Equal(t, bm.Test(10), false, "bm.Test(10) != false")
	bm.Set(10)
	bm.Test(10)
	assert.Equal(t, bm.Test(10), true, "bm.Test(10) != true")
	assert.Equal(t, bm.Output(size), "0000000000100000", "bm.Output(size) != 0000000000100000")

	bm.Clear(10)
	assert.Equal(t, bm.Test(10), false, "bm.Test(10) != false")
	assert.Equal(t, bm.Output(size), "0000000000000000", "bm.Output(size) != 0000000000000000")

	bm.Set(size + 2)
	assert.Equal(t, bm.Test(size+2), true, "bm.Test(size+2) != true")
}
