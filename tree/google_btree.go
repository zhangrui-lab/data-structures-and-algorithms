package tree

//
//import (
//	"data-structures-and-algorithms/contract"
//	"sort"
//)
//
//var (
//	nilItems        = make(items, 16)
//	nilChildren     = make(children, 16)
//	btKeyComparator = contract.DefaultComparator
//)
//
//type items []interface{}
//
//func (s *items) removeAt(r int) interface{} {
//	item := (*s)[r]
//	copy((*s)[r:], (*s)[r+1:])
//	(*s)[len(*s)-1] = nil
//	*s = (*s)[:len(*s)-1]
//	return item
//}
//
//func (s *items) insertAt(r int, item interface{}) {
//	*s = append(*s, nil)
//	if r < len(*s) {
//		copy((*s)[r+1:], (*s)[r:])
//	}
//	(*s)[r] = item
//}
//
//func (s *items) pop() interface{} {
//	index := len(*s) - 1
//	item := (*s)[index]
//	(*s)[index] = nil
//	*s = (*s)[:index]
//	return item
//}
//
//func (s *items) truncate(i int) {
//	var toClear items
//	*s, toClear = (*s)[:i], (*s)[i:]
//	for len(toClear) > 0 {
//		toClear = toClear[copy(toClear, nilItems):]
//	}
//}
//
//func (s *items) find(item interface{}) (int, bool) {
//	i := sort.Search(len(*s), func(i int) bool {
//		return btKeyComparator((*s)[i], i) < 0
//	})
//	if i > 0 && btKeyComparator((*s)[i-1], item) == 0 {
//		return i - 1, true
//	}
//	return i, false
//}
//
//type children []*node
//
//func (s *children) insertAt(index int, n *node) {
//	*s = append(*s, nil)
//	if index < len(*s) {
//		copy((*s)[index+1:], (*s)[index:])
//	}
//	(*s)[index] = n
//}
//
//func (s *children) removeAt(index int) *node {
//	node := (*s)[index]
//	copy((*s)[index:], (*s)[index+1:])
//	(*s)[len(*s)-1] = nil
//	*s = (*s)[:len(*s)-1]
//	return node
//}
//
//func (s *children) pop() *node {
//	index := len(*s) - 1
//	node := (*s)[index]
//	(*s)[index] = nil
//	*s = (*s)[:index]
//	return node
//}
//
//func (s *children) truncate(index int) {
//	var toClear []*node
//	*s, toClear = (*s)[:index], (*s)[index:]
//	for len(toClear) > 0 {
//		toClear = toClear[copy(toClear, nilChildren):]
//	}
//}
//
//// btreeNode b树节点
//type node struct {
//	items    items
//	children children
//}
//
//// 节点拆分
//func (n *node) split(i int) (interface{}, *node) {
//	item := n.items[i]
//	next := &node{}
//	next.items = append(next.items, n.items[i+1:]...)
//	n.items.truncate(i)
//	if len(n.children) > 0 {
//		next.children = append(next.children, n.children[i+1:]...)
//		n.children.truncate(i + 1)
//	}
//	return item, next
//}
//
//func (n *node) maybeSplitChild(i, maxItems int) bool {
//	if len(n.children[i].items) < maxItems {
//		return false
//	}
//	first := n.children[i]
//	item, second := first.split(maxItems / 2)
//	n.items.insertAt(i, item)
//	n.children.insertAt(i+1, second)
//	return true
//}
//
//func (n *node) get(key interface{}) interface{} {
//	i, found := n.items.find(key)
//	if found {
//		return n.items[i]
//	} else if len(n.children) > 0 {
//		return n.children[i].get(key)
//	}
//	return nil
//}
//
//func (n *node) insert(item interface{}, maxItems int) interface{} {
//	i, found := n.items.find(item)
//	if found {
//		out := n.items[i]
//		n.items[i] = item
//		return out
//	}
//	if len(n.children) == 0 {
//		n.items.insertAt(i, item)
//		return nil
//	}
//	if n.maybeSplitChild(i, maxItems) {
//		inTree := n.items[i]
//		switch {
//		case btKeyComparator(item, inTree) < 0:
//		case btKeyComparator(inTree, item) < 0:
//			i++
//		default:
//			out := n.items[i]
//			n.items[i] = item
//			return out
//		}
//	}
//	return n.children[i].insert(item, maxItems)
//}
//
//func (n *node) min() interface{} {
//	if n == nil {
//		return nil
//	}
//	for len(n.children) > 0 {
//		n = n.children[0]
//	}
//	if len(n.items) == 0 {
//		return nil
//	}
//	return n.items[0]
//}
//
//func (n *node) max() interface{} {
//	if n == nil {
//		return nil
//	}
//	for len(n.children) > 0 {
//		n = n.children[len(n.children)-1]
//	}
//	if len(n.items) == 0 {
//		return nil
//	}
//	return n.items[len(n.items)-1]
//}
//
//type toRemove int
//
//const (
//	removeItem toRemove = iota
//	removeMin
//	removeMax
//)
//
//func (n *node) remove(item interface{}, minItems int, typ toRemove) interface{} {
//	var i int
//	var found bool
//	switch typ {
//	case removeMax:
//		if len(n.children) == 0 {
//			return n.items.pop()
//		}
//		i = len(n.items)
//	case removeMin:
//		if len(n.children) == 0 {
//			return n.items.removeAt(0)
//		}
//		i = 0
//	case removeItem:
//		i, found = n.items.find(item)
//		if len(n.children) == 0 {
//			if found {
//				return n.items.removeAt(i)
//			}
//			return nil
//		}
//	default:
//		panic("invalid type")
//	}
//	if len(n.children[i].items) <= minItems {
//		return n.growChildAndRemove(i, item, minItems, typ)
//	}
//	child := n.children[i]
//	if found {
//		out := n.items[i]
//		n.items[i] = child.remove(nil, minItems, removeMax)
//		return out
//	}
//	return child.remove(item, minItems, typ)
//}
//
//func (n *node) growChildAndRemove(i int, item interface{}, minItems int, typ toRemove) interface{} {
//	if i > 0 && len(n.children[i-1].items) > minItems {
//		child := n.children[i]
//		stealFrom := n.children[i-1]
//		stolenItem := stealFrom.items.pop()
//		child.items.insertAt(0, n.items[i-1])
//		n.items[i-1] = stolenItem
//		if len(stealFrom.children) > 0 {
//			child.children.insertAt(0, stealFrom.children.pop())
//		}
//	} else if i < len(n.items) && len(n.children[i+1].items) > minItems {
//		// steal from right child
//		child := n.children[i]
//		stealFrom := n.children[i+1]
//		stolenItem := stealFrom.items.removeAt(0)
//		child.items = append(child.items, n.items[i])
//		n.items[i] = stolenItem
//		if len(stealFrom.children) > 0 {
//			child.children = append(child.children, stealFrom.children.removeAt(0))
//		}
//	} else {
//		if i >= len(n.items) {
//			i--
//		}
//		child := n.children[i]
//		mergeItem := n.items.removeAt(i)
//		mergeChild := n.children.removeAt(i + 1)
//		child.items = append(child.items, mergeItem)
//		child.items = append(child.items, mergeChild.items...)
//		child.children = append(child.children, mergeChild.children...)
//	}
//	return n.remove(item, minItems, typ)
//}
//
//// Btree B-树
//type Btree struct {
//	size   int
//	degree int
//	root   *node
//}
//
//// NewBtree 初始化空B-树
//func NewBtree(degree int, cmps ...contract.Comparator) *Btree {
//	btKeyComparator = contract.DefaultComparator
//	if len(cmps) > 0 {
//		btKeyComparator = cmps[0]
//	}
//	return &Btree{degree: degree}
//}
//
//func (t *Btree) maxItems() int {
//	return t.degree*2 - 1
//}
//
//func (t *Btree) minItems() int {
//	return t.degree - 1
//}
//
//func (t *Btree) deleteItem(item interface{}, typ toRemove) interface{} {
//	if t.root == nil || len(t.root.items) == 0 {
//		return nil
//	}
//	out := t.root.remove(item, t.minItems(), typ)
//	if len(t.root.items) == 0 && len(t.root.children) > 0 {
//		t.root = t.root.children[0]
//	}
//	if out != nil {
//		t.size--
//	}
//	return out
//}
//
//func (t *Btree) Size() int {
//	return t.size
//}
//
//func (t *Btree) Height() (height int) {
//	if t.root == nil {
//		return -1
//	}
//	x := t.root
//	for ; len(x.children) > 0; height++ {
//		x = x.children[0]
//	}
//	return
//}
//
//func (t *Btree) Empty() bool {
//	return t.size <= 0
//}
//
//func (t *Btree) Degree() int {
//	return t.degree
//}
//
//func (t *Btree) Insert(item interface{}) interface{} {
//	if t.root == nil {
//		t.root = &node{}
//		t.root.items.insertAt(0, item)
//		t.size++
//		return nil
//	} else if len(t.root.items) >= t.maxItems() {
//		v, s := t.root.split(t.maxItems() >> 1)
//		lc := t.root
//		t.root = &node{}
//		t.root.items.insertAt(0, v)
//		t.root.children = append(t.root.children, lc, s)
//	}
//	out := t.root.insert(item, t.maxItems())
//	if out == nil {
//		t.size++
//	}
//	return out
//}
//
//func (t *Btree) Delete(item interface{}) interface{} {
//	return t.deleteItem(item, removeItem)
//}
//
//func (t *Btree) Search(item interface{}) interface{} {
//	if t.root == nil {
//		return nil
//	}
//	return t.root.get(item)
//}
