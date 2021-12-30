// Package tree 前缀树节点信息：可用于 Trie树，Radix树等树节点
package tree

import "sort"

// WalkFn 对节点进行遍历：返回true时终止迭代
type WalkFn func(key string, value interface{}) bool

var _ sort.Interface = children(nil)

// 树节点
type innerNode struct {
	char     byte
	children children
	leaf     *leafNode
}

// 值
type leafNode struct {
	key   string
	value interface{}
}

// 子节点信息
type children []*innerNode

func newInnerNode(char byte) *innerNode {
	return &innerNode{char: char, children: make([]*innerNode, 0), leaf: nil}
}

// Len 子节点排序支持
func (c children) Len() int {
	return len(c)
}

// Less 子节点排序支持
func (c children) Less(i, j int) bool {
	return c[i].char < c[j].char
}

// Swap 子节点排序支持
func (c children) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// 插入子节点
func (inner *innerNode) insertChild(char byte) *innerNode {
	c := newInnerNode(char)
	inner.children = append(inner.children, c)
	sort.Sort(inner.children)
	return c
}

// 插入值
func (inner *innerNode) insertLeaf(key string, value interface{}) {
	inner.leaf = &leafNode{key: key, value: value}
}

// 获取子节点
func (inner *innerNode) getChild(char byte) *innerNode {
	index := sort.Search(len(inner.children), func(i int) bool {
		return inner.children[i].char >= char
	})
	if index < len(inner.children) && inner.children[index].char == char {
		return inner.children[index]
	}
	return nil
}

// 删除子节点
func (inner *innerNode) delChild(char byte) {
	index := sort.Search(len(inner.children), func(i int) bool {
		return inner.children[i].char >= char
	})
	if index < len(inner.children) && inner.children[index].char == char {
		copy(inner.children[index:], inner.children[index+1:])
		inner.children[len(inner.children)-1] = nil
		inner.children = inner.children[:len(inner.children)-1]
	}

}

// 以当前节点为根对跟定key进行查找
func (inner *innerNode) search(key string) *innerNode {
	for i := 0; i < len(key) && inner != nil; i = i + 1 {
		inner = inner.getChild(key[i])
	}
	return inner
}

// 是否存在值
func (inner *innerNode) isLeaf() bool {
	return inner.leaf != nil
}

// 对以当前节点为根的子树进行dfs遍历
func (inner *innerNode) dfs(call WalkFn) bool {
	if inner.isLeaf() && call(inner.leaf.key, inner.leaf.value) {
		return true
	}
	for _, c := range inner.children {
		if c.dfs(call) {
			return true
		}
	}
	return false
}
