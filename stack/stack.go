// Package stack 堆栈
package stack

// Stack 栈
type Stack struct {
	elem []interface{}
}

// New 初始化空栈
func New() *Stack {
	return new(Stack)
}

// Empty 当前栈是否为空
func (s *Stack) Empty() bool {
	if s.elem == nil || len(s.elem) == 0 {
		return true
	}
	return false
}

// Push 入栈
func (s *Stack) Push(e interface{}) {
	//if s.elem == nil {
	//	s.elem = make([]interface{}, 0, 3)
	//}
	s.elem = append(s.elem, e)
}

// Pop 出栈：栈为空时返回nil且不做操作
func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	e := s.elem[len(s.elem)-1]
	s.elem = s.elem[:len(s.elem)-1]
	return e
}

// Top 栈顶元素信息：栈为空时返回nil
func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	return s.elem[len(s.elem)-1]
}
