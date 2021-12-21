// Package tree 二叉树：仅提供树结构的基础操作支持，不包含指定元素值 v 的插入与查找等（查找树中实现）
package tree

import (
	"data-structures-and-algorithms/contract"
)

// BinTree 二叉树
type BinTree struct {
	size int
	root *BinNode
}

// New 创建空树
func New() *BinTree {
	return &BinTree{}
}

// Size 树规模
func (tree *BinTree) Size() int {
	return tree.size
}

// Empty 是否为空树
func (tree *BinTree) Empty() bool {
	return tree.size <= 0
}

// Root 树根
func (tree *BinTree) Root() *BinNode {
	return tree.root
}

// InsertAsRoot 插入根节点
func (tree *BinTree) InsertAsRoot(key interface{}) *BinNode {
	tree.size = 1
	tree.root = newBinNode(key, key, nil, nil, nil)
	return tree.root
}

// InsertAsLc e作为x的左孩子（原无）插入
func (tree *BinTree) InsertAsLc(x *BinNode, key interface{}) *BinNode {
	tree.size++
	x.lc = newBinNode(key, key, x, nil, nil)
	return x.lc
}

// InsertAsRc e作为x的右孩子（原无）插入
func (tree *BinTree) InsertAsRc(x *BinNode, key interface{}) *BinNode {
	tree.size++
	x.rc = newBinNode(key, key, x, nil, nil)
	return x.rc
}

// AttachAsLC 二叉树子树接入算法：将S当作节点x的左子树 (tree无左子树) 接入，S本身置空，并返回接入位置
func (tree *BinTree) AttachAsLC(other *BinTree) *BinNode {
	tree.root.lc = other.root
	if tree.root.lc != nil {
		tree.root.lc.parent = tree.root
	}
	tree.size += other.size
	other.size = 0
	other.root = nil
	return tree.root.lc
}

// AttachAsRC other作为x右子树接入
func (tree *BinTree) AttachAsRC(other *BinTree) *BinNode {
	tree.root.rc = other.root
	if tree.root.rc != nil {
		tree.root.rc.parent = tree.root
	}
	tree.size += other.size
	other.size = 0
	other.root = nil
	return tree.root.rc
}

// Remove 删除二叉树中位置 x （x为二叉树中的合法位置）处的节点及其后代，返回被删除节点的数值
func (tree *BinTree) Remove(x *BinNode) int {
	return tree.removeAt(x)
}

// Secede 将子树 x（x为二叉树中的合法位置）从当前树中摘除，并将其转换为一棵独立子树
func (tree *BinTree) Secede(x *BinNode) *BinTree {
	return &BinTree{root: x, size: tree.removeAt(x)}
}

// TravelLevel 层次遍历
func (tree *BinTree) TravelLevel(visitor contract.KvVisitor) {
	tree.root.travelLevel(visitor)
}

// TravelPre 先序遍历
func (tree *BinTree) TravelPre(visitor contract.KvVisitor) {
	tree.root.travelPre(visitor)
}

// TravelIn 中序遍历
func (tree *BinTree) TravelIn(visitor contract.KvVisitor) {
	tree.root.travelIn(visitor)
}

// TravelPost 后序遍历
func (tree *BinTree) TravelPost(visitor contract.KvVisitor) {
	tree.root.travelPost(visitor)
}

// 从tree中移除节点x
func (tree *BinTree) removeAt(x *BinNode) int {
	if x == nil {
		return 0
	}
	if x.isRoot() {
		tree.root = nil
		size := tree.size
		tree.size = 0
		return size
	}
	if x.isLc() {
		x.parent.lc = nil
	} else {
		x.parent.rc = nil
	}
	x.parent = nil
	tree.size -= x.size()
	return x.size()
}
