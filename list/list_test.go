package list

import (
	"data-structures-and-algorithms/types"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 长度验证
func checkSize(t *testing.T, list *List, n int) {
	assert.Equal(t, list.Size(), n, fmt.Errorf("l.Size() = %d, want %d", list.size, n))
}

// 空列表首位哨兵验证
func checkEmpty(t *testing.T, list *List) {
	err := fmt.Errorf("l.header.prev = %p, l.header.next = %p, l.trailer.prev = %p,  l.trailer.next = %p;", list.header.prev, list.header.next, list.trailer.prev, list.trailer.next)
	assert.Equal(t, list.header.prev, (*Node)(nil), err)
	assert.Equal(t, list.header.next, list.trailer, err)

	assert.Equal(t, list.trailer.next, (*Node)(nil), err)
	assert.Equal(t, list.trailer.prev, list.header, err)

	assert.Equal(t, list.Size(), 0, "empty list size() != 0")
	assert.Equal(t, list.Empty(), true, "empty list empty() != true")
}

// 验证e1为e2直接后继
func checkIsNext(t *testing.T, e1, e2 *Node) {
	checkIsPrev(t, e2, e1)
}

// 验证e1为e2直接前驱
func checkIsPrev(t *testing.T, e1, e2 *Node) {
	err := "e1.next != e2 or e2.prev != e1"
	assert.Equal(t, e1.next, e2, err)
	assert.Equal(t, e2.prev, e1, err)
}

func TestEmpty(t *testing.T) {
	l := New()
	checkEmpty(t, l)

	l.PushBack(types.Int(10))
	l.PushBack(types.Int(20))
	checkSize(t, l, 2)

	l.Clear()
	checkEmpty(t, l)
}

func TestGet(t *testing.T) {
	v1, v2, v3, v4 := types.Int(1), types.Int(2), types.Int(3), types.Int(4)
	l := New()
	l.PushFront(v1)
	l.PushFront(v2)
	l.PushFront(v3)
	l.PushFront(v4) //4->3->2->1
	assert.Equal(t, l.String(), "{4, 3, 2, 1}")

	checkSize(t, l, 4)
	e := l.At(1)
	assert.NotEqual(t, e, nil, "l.At(1) == nil")
	assert.Equal(t, e.Data, v3, "l.At(1).Data != v3")

	e = l.Front()
	assert.NotEqual(t, e, nil, "l.Front() == nil")
	assert.Equal(t, e.Data, v4, "l.Front() != v4")

	e = l.Back()
	assert.NotEqual(t, e, nil, "l.Back() == nil")
	assert.Equal(t, e.Data, v1, "l.Back() != v1")

	v := l.Remove(l.Back())
	assert.Equal(t, v, v1, "l.Remove(l.Back()) != v1")

	v = l.Remove(l.Front())
	assert.Equal(t, v, v4, "l.Remove(l.Front()) != v4")

	assert.Equal(t, l.String(), "{3, 2}") // 3-2
	checkSize(t, l, 2)

	e = l.At(1)
	assert.NotEqual(t, e, nil, "after delete front,back. l.At(1) == nil")
	assert.Equal(t, e.Data, v2, "after delete front,back. l.At(2).Data != v2")

	e = l.Front()
	assert.NotEqual(t, e, nil, "after delete front,back. l.Front() == nil")
	assert.Equal(t, e.Data, v3, "after delete front,back. l.Front() != v3")

	e = l.Back()
	assert.NotEqual(t, e, nil, "after delete front,back. l.Back() == nil")
	assert.Equal(t, e.Data, v2, "after delete front,back. l.Back() != v2")
}

func TestInsertAndRemove(t *testing.T) {
	v1, v2, v3, v4, v5 := types.Int(1), types.Int(2), types.Int(3), types.Int(4), types.Int(5)
	l := New()

	// -- insert

	e1 := l.PushBack(v1) // v1
	checkIsPrev(t, e1, l.trailer)
	checkIsNext(t, e1, l.header)

	e2 := l.PushFront(v2) // v2->v1
	checkIsPrev(t, e2, e1)
	checkIsNext(t, e2, l.header)

	e3 := l.Insert(v3, e2) // v2->v3->v1
	checkIsPrev(t, e3, e1)
	checkIsNext(t, e3, e2)

	e4 := l.InsertBefore(v4, e3) // v2->v4->v3->v1
	checkIsPrev(t, e4, e3)
	checkIsNext(t, e4, e2)

	e5 := l.InsertAfter(v5, e3) // v2->v4->v3->v5->v1
	checkIsPrev(t, e5, e1)
	checkIsNext(t, e5, e3)

	checkSize(t, l, 5)
	assert.Equal(t, l.String(), "{2, 4, 3, 5, 1}", "l.String() != {2, 4, 3, 5, 1}")

	// --- remove		v2->v4->v3->v5->v1

	v := l.PopBack() // v2->v4->v3->v5
	assert.Equal(t, v, v1, "l.PopBack() != v1")
	checkIsPrev(t, e5, l.trailer)

	v = l.PopFront() // v4->v3->v5
	assert.Equal(t, v, v2, "l.PopFront() != v2")
	checkIsNext(t, e4, l.header)

	v = l.Remove(e2) // v4->v3->v5
	assert.Equal(t, v, nil, "l.Remove(e2) != 0 (removed)")

	checkSize(t, l, 3)
	assert.Equal(t, l.String(), "{4, 3, 5}", "l.String() != {4, 3, 5}")
}

func TestMove(t *testing.T) {
	v1, v2, v3, v4, v5 := types.Int(1), types.Int(2), types.Int(3), types.Int(4), types.Int(5)
	l := New()
	e1 := l.PushFront(v1)
	e2 := l.PushFront(v2)
	e3 := l.PushFront(v3)
	e4 := l.PushFront(v4)
	e5 := l.PushFront(v5) // {5, 4, 3, 2, 1}

	checkSize(t, l, 5)
	assert.Equal(t, l.String(), "{5, 4, 3, 2, 1}", "l.String() != {5, 4, 3, 2, 1}")
	assert.Equal(t, l.disordered(), 4, "disordered {5, 4, 3, 2, 1} != 4")

	l.MoveAfter(e2, e1) // {5, 4, 3, 1, 2}
	checkIsPrev(t, e2, l.trailer)
	checkIsNext(t, e2, e1)
	checkIsNext(t, e1, e3)

	l.MoveBefore(e4, e5) // {4, 5, 3, 1, 2}
	checkIsPrev(t, e4, e5)
	checkIsNext(t, e4, l.header)
	checkIsPrev(t, e5, e3)

	l.MoveToBack(e3) // {4, 5, 1, 2, 3}
	checkIsPrev(t, e3, l.trailer)
	checkIsNext(t, e3, e2)
	checkIsPrev(t, e5, e1)

	assert.Equal(t, l.String(), "{4, 5, 1, 2, 3}", "l.String() != {4, 5, 1, 2, 3}")
	assert.Equal(t, l.disordered(), 1, "disordered {4, 5, 1, 2, 3} != 1")
}

func TestMerge(t *testing.T) {
	v1, v2, v3, v4, v5 := types.Int(1), types.Int(2), types.Int(3), types.Int(4), types.Int(5)
	l1 := New()
	l1.PushFront(v1)
	l1.PushFront(v2) // {2, 1}
	assert.Equal(t, l1.String(), "{2, 1}", "l1.String() != {2, 1}")

	l2 := New()
	l2.PushBack(v3)
	l2.PushBack(v5)
	l2.PushBack(v4) // {3, 5, 4}
	assert.Equal(t, l2.String(), "{3, 5, 4}", "l2.String() != {3, 5, 4}")

	l1.PushBackList(l2) // {2, 1, 3, 5, 4}
	assert.Equal(t, l1.String(), "{2, 1, 3, 5, 4}", "l1.String() != {2, 1, 3, 5, 4}")

	l2.PushFrontList(l1) // {2, 1, 3, 5, 4, 3, 5, 4}
	assert.Equal(t, l2.String(), "{2, 1, 3, 5, 4, 3, 5, 4}", "l2.String() != {2, 1, 3, 5, 4, 3, 5, 4}")

}

func TestReverse(t *testing.T) {
	v1, v2, v3, v4, v5 := types.Int(1), types.Int(2), types.Int(3), types.Int(4), types.Int(5)
	l1 := New()
	l1.PushFront(v1)
	l1.PushFront(v2) // {2, 1}
	assert.Equal(t, l1.String(), "{2, 1}", "l1.String() != {2, 1}")
	l1.reverse() // {1, 2}
	assert.Equal(t, l1.String(), "{1, 2}", "l1.String() != {1, 2}")

	l2 := New()
	l2.PushBack(v3)
	l2.PushBack(v5)
	l2.PushBack(v4) // {3, 5, 4}
	assert.Equal(t, l2.String(), "{3, 5, 4}", "l2.String() != {3, 5, 4}")
}
