package skiplist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSkipList(t *testing.T) {
	k1, k2, k3, k4, k5, k6, k7, k8 := 1, 2, 3, 4, 5, 6, 7, 8
	v1, v2, v3, v4, v5, v6, v7, v8 := 10, 20, 30, 40, 50, 60, 70, 80
	skl := NewSkipList()
	skl.Insert(k5, v5)
	skl.Insert(k2, v2)
	skl.Insert(k1, v1)
	skl.Insert(k4, v4)
	skl.Insert(k7, v7)
	skl.Insert(k3, v3)
	skl.Insert(k8, v8)
	skl.Insert(k6, v6)
	fmt.Println()
	//{1}
	//{1,2,3,4}
	//{1,2,3,4,5,6,7,8}

	assert.Equal(t, skl.Search(k1), v1, "skl.Search(k1) != v1")
	assert.Equal(t, skl.Search(k2), v2, "skl.Search(k2) != v2")
	assert.Equal(t, skl.Size(), 8, "skl.Size() != 8")

	skl.Remove(k1)
	skl.Remove(k6)
	skl.Remove(k3)
	skl.Remove(k8)
	//{2,4}
	//{2,4,5,7}

}
