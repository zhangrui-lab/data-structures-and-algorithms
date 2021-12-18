package forward_list

import (
	"data-structures-and-algorithms/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardList(t *testing.T) {
	fl := New()
	assert.Equal(t, fl.Size(), 0, "empty forwardList size() != 0")
	assert.Equal(t, fl.Empty(), true, "empty forwardList Empty() != true")

	v1, v2, v3 := types.Int(1), types.Int(2), types.Int(3)
	e1 := fl.PushFront(v1) // 1
	assert.Equal(t, e1.Data, v1, "e1.Data != v1")
	assert.Equal(t, fl.Size(), 1, "size() != 1")

	e2 := fl.PushFront(v2) // 2->1
	assert.Equal(t, e2.next.Data, v1, "e2.next.Data != v1")
	assert.Equal(t, fl.Size(), 2, "size() != 2")

	assert.Equal(t, fl.String(), "{2, 1}", "forwardList.String() != {2, 1}")

	v := fl.RemoveAfter(e2) // 2
	assert.Equal(t, v, v1, "fl.RemoveAfter(e2) != v1")

	v = fl.RemoveAfter(e2) // 2
	assert.Equal(t, v, nil, "e2.Data != nil")
	assert.Equal(t, fl.Size(), 1, "size() != 1")

	fl.PushFront(v2)
	fl.PushFront(v2)
	fl.PushFront(v1)
	fl.PushFront(v3) // 3->1->2->2->2
	assert.Equal(t, fl.String(), "{3, 1, 2, 2, 2}", "forwardList.String() != {3, 1, 2, 2, 2}")

	cv := fl.Remove(v2) // 3->1
	assert.Equal(t, cv, 3, "Remove(v2) != 3")

	assert.Equal(t, fl.String(), "{3, 1}", "forwardList.String() != {1}")

	v = fl.PopFront() // 1
	assert.Equal(t, v, v3, "fl.PopFront() != v3")

	cv = fl.Clear()
	assert.Equal(t, cv, 1, "Clear() != 1")
	assert.Equal(t, fl.Size(), 0, "cleared forwardList.Size() != 0")
	assert.Equal(t, fl.Empty(), true, "cleared forwardList.Empty() != true")
}
