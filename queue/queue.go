// Package queue 队列
package queue

// Queue 队列
type Queue struct {
	elem []interface{}
}

// New 初始化空栈
func New() *Queue {
	return new(Queue)
}

// Empty 当前队列是否为空
func (q *Queue) Empty() bool {
	return len(q.elem) <= 0
}

// Size 队列元素个数
func (q *Queue) Size() int {
	return len(q.elem)
}

// Front 队首元素信息：空队列返回nil
func (q *Queue) Front() interface{} {
	if q.Empty() {
		return nil
	}
	return q.elem[0]
}

// Back 队尾元素信息：空队列返回nil
func (q *Queue) Back() interface{} {
	if q.Empty() {
		return nil
	}
	return q.elem[len(q.elem)-1]
}

// Push 入队
func (q *Queue) Push(e interface{}) {
	q.elem = append(q.elem, e)
}

// Pop 出对：空队列不做操作并返回nil
func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	e := q.elem[0]
	q.elem = q.elem[1:]
	return e
}
