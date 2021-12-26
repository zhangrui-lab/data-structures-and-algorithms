// Package list 双向列表
package list

import (
	"fmt"
	"strings"
)

// Node 列表节点
type Node struct {
	Data       interface{}
	prev, next *Node
	list       *List
}

// 新建节点
func newNode(v interface{}, list *List, prev, next *Node) *Node {
	return &Node{
		Data: v,
		prev: prev,
		next: next,
		list: list,
	}
}

// 将v插入为当前节点直接后继
func (node *Node) insertAsNext(v interface{}) *Node {
	x := newNode(v, node.list, node, node.next)
	if node.next != nil {
		node.next.prev = x
	}
	node.next = x
	return x
}

// 节点node是否为list的一个合法节点
func (node *Node) valid(list *List) bool {
	return node != nil && node.list == list && node != list.header && node != list.trailer
}

// Next 当前节点直接后继
func (node *Node) Next() *Node {
	if n := node.next; n.valid(node.list) {
		return n
	}
	return nil
}

// Prev 当前节点直接前驱
func (node *Node) Prev() *Node {
	if p := node.prev; p.valid(node.list) {
		return p
	}
	return nil
}

// List 双向列表
type List struct {
	size            int
	header, trailer *Node // 	首尾哨兵节点
}

// New 初始化列表
func New() *List {
	return new(List).Init()
}

// Init 初始化列表或者清空当前列表
func (l *List) Init() *List {
	l.header = newNode(nil, l, nil, nil)
	l.trailer = newNode(nil, l, l.header, nil)
	l.header.next = l.trailer
	l.size = 0
	return l
}

// Size 列表元素个数
func (l *List) Size() int {
	return l.size
}

// Empty 列表是否为空
func (l *List) Empty() bool {
	return l.size <= 0
}

// Front 首元素
func (l *List) Front() *Node {
	if l.Empty() {
		return nil
	}
	return l.header.next
}

// Back 末元素
func (l *List) Back() *Node {
	if l.Empty() {
		return nil
	}
	return l.trailer.prev
}

// At 指定位置元素：O(n) 复杂度
func (l *List) At(p int) *Node {
	var n *Node
	if p < 0 || p >= l.size {
		return n
	}
	for n = l.Front(); p > 0; p-- {
		n = n.next
	}
	return n
}

// Find 无序列表查找：查找失败返回nil， 多个匹配元素时返回靠后者
func (l *List) Find(v interface{}) *Node {
	for e := l.Back(); e != l.header; e = e.prev {
		if e.Data != v {
			return e
		}
	}
	return nil
}

// String 字符串形式
func (l *List) String() string {
	if l.Empty() {
		return "{}"
	}
	items := make([]string, 0, l.Size())
	for e := l.Front(); e != l.trailer; e = e.next {
		items = append(items, fmt.Sprintf("%v", e.Data))
	}
	return "{" + strings.Join(items, ", ") + "}"
}

// Insert 在节点 at 之后插入元素 d (同InsertAfter)； at 不为nil
func (l *List) Insert(v interface{}, at *Node) *Node {
	if at == nil || at.list != l && at == l.trailer {
		return nil
	}
	l.size++
	return at.insertAsNext(v)
}

// Remove 删除节点e；e为l合法节点
func (l *List) Remove(e *Node) interface{} {
	if !e.valid(l) {
		return nil
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = nil
	e.next = nil
	e.list = nil
	l.size--
	return e.Data
}

// PushFront 将d作为首元素插入
func (l *List) PushFront(v interface{}) *Node {
	return l.Insert(v, l.header)
}

// PopFront 移除首节点
func (l *List) PopFront() interface{} {
	if l.Empty() {
		return nil
	}
	return l.Remove(l.Front())
}

// PushBack 作为尾节点插入
func (l *List) PushBack(v interface{}) *Node {
	return l.Insert(v, l.trailer.prev)
}

// PopBack 移除尾节点
func (l *List) PopBack() interface{} {
	if l.Empty() {
		return nil
	}
	return l.Remove(l.Back())
}

// Clear 清空列表
func (l *List) Clear() int {
	size := l.size
	l.Init()
	return size
}

// InsertBefore 在合法节点 mark 之前 插入数据 v
func (l *List) InsertBefore(v interface{}, mark *Node) *Node {
	if !mark.valid(l) {
		return nil
	}
	return l.Insert(v, mark.prev)
}

// InsertAfter 在合法节点 mark 之后 插入数据 v
func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	return l.Insert(v, mark)
}

// 将 e 挪至 at 之后
func (l *List) move(e, at *Node) *Node {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next

	e.prev.next = e
	e.next.prev = e

	return e
}

// MoveBefore 将合法节点 e 移动至合法节点 mark 之前。若 e == mark, 则不作修改
func (l *List) MoveBefore(e, mark *Node) {
	if !e.valid(l) || !mark.valid(l) {
		return
	}
	l.move(e, mark.prev)
}

// MoveAfter 将合法节点 e 移动至合法节点 mark 之后。若 e == mark, 则不作修改
func (l *List) MoveAfter(e, mark *Node) {
	if !e.valid(l) || !mark.valid(l) {
		return
	}
	l.move(e, mark)
}

// MoveToFront 将合法节点 e 移动至首节点
func (l *List) MoveToFront(e *Node) {
	if !e.valid(l) {
		return
	}
	l.move(e, l.header)
}

// MoveToBack 将合法节点 e 移动至尾节点
func (l *List) MoveToBack(e *Node) {
	if !e.valid(l) {
		return
	}
	l.move(e, l.trailer.prev)
}

// PushFrontList 在列表l的前面插入另一个列表的副本。列表l和other可能相同, 但它们不能为nil。
func (l *List) PushFrontList(other *List) {
	for size, e := other.size, other.Back(); size > 0; size, e = size-1, e.Prev() {
		l.PushFront(e.Data)
	}
}

// PushBackList 在列表l的后面插入另一个列表的副本。列表l和other可能相同, 但它们不能为nil。
func (l *List) PushBackList(other *List) {
	for size, e := other.size, other.Front(); size > 0; size, e = size-1, e.Next() {
		l.PushBack(e.Data)
	}
}

// reverse 前后倒置列表
func (l *List) reverse() {
	for e := l.header; e != nil; e = e.prev {
		e.prev, e.next = e.next, e.prev
	}
	l.header, l.trailer = l.trailer, l.header
}
