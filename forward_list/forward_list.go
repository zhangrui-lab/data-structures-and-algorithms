package forward_list

import (
	"fmt"
	"strings"
)

// Node 单向列表节点
type Node struct {
	Data interface{}
	next *Node
}

// insert 插入v为当前节点后继
func (e *Node) insert(v interface{}) *Node {
	p := &Node{Data: v, next: e.next}
	e.next = p
	return p
}

// ForwardList 单向列表
type ForwardList struct {
	size int
	root *Node
}

// New 初始化列表
func New() *ForwardList {
	return &ForwardList{
		size: 0,
		root: &Node{Data: nil, next: nil},
	}
}

// Size 列表元素个数
func (l *ForwardList) Size() int {
	return l.size
}

// Empty 列表是否为空
func (l *ForwardList) Empty() bool {
	return l.size <= 0
}

// Front 首元素
func (l *ForwardList) Front() *Node {
	return l.root.next
}

// RemoveAfter at的直接后继存在时，将其移除并返回其值
func (l *ForwardList) RemoveAfter(at *Node) interface{} {
	e := at.next
	if e == nil {
		return nil
	}
	l.size--
	at.next = e.next
	return e.Data
}

// PushFront 将v作为首元素值插入
func (l *ForwardList) PushFront(v interface{}) *Node {
	l.size++
	return l.root.insert(v)
}

// PopFront 非空时移除首元素并返回其值
func (l *ForwardList) PopFront() interface{} {
	return l.RemoveAfter(l.root)
}

// Clear 清空单向列表, 并返回删除元素个数
func (l *ForwardList) Clear() int {
	l.root.next = nil
	size := l.size
	l.size = 0
	return size
}

// Merge 将other元素追加到l末尾
func (l *ForwardList) Merge(other *ForwardList) int {
	end := l.root
	for ; end.next != nil; end = end.next {
	}
	for e := other.root.next; e != nil; e = e.next {
		end = end.insert(e.Data)
	}
	l.size += other.size
	return other.size
}

// Remove 将值为v的元素从l中移除
func (l *ForwardList) Remove(v interface{}) int {
	size := l.size
	for e := l.root; e.next != nil; {
		if e.next.Data != v {
			e = e.next
			continue
		}
		l.RemoveAfter(e)
	}
	return size - l.size
}

// String 字符串形式
func (l *ForwardList) String() string {
	if l.Empty() {
		return "{}"
	}
	items := make([]string, 0, l.Size())
	for e := l.root; e.next != nil; e = e.next {
		items = append(items, fmt.Sprintf("%v", e.next.Data))
	}
	return "{" + strings.Join(items, ", ") + "}"
}
