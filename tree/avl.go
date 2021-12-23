// Package tree AVL树，基于二叉查找树的基础上，新增局部节点施加约束：任意节点的左右子树高度差(平衡因子)不大于2.
package tree

import (
	"data-structures-and-algorithms/contract"
)

// Avl Avl树
type Avl struct {
	Bst
}

// NewAvl 新建空avl树
func NewAvl(cmps ...contract.Comparator) *Avl {
	return &Avl{Bst: *NewBst(cmps...)}
}

// Insert Avl树插入
// 节点的插入引发的最低失衡节点必定不会为其父节点。
// 节点的插入引发的失衡在一次调整之后即可回复，不会向上传递。
// 恢复算法：
// 1. 从 t.hot 开始，找到最低失衡节点 g。
// 2. 若 g 存在， 则考虑 g 在 x 后代节点侧的子节点 p， 孙子节点 v。
// 3. 依据 g， p， x 之间的相对位置做节点旋转。
// 3. 依据 g， p， x 之间的相对位置做节点旋转。
// 3.1. p.isLc && x.isLc: 对 g 做一次顺时针旋转即可。
// 3.2. p.isLc && x.isRc: 对 p 做一次逆时针旋转， g 做一次顺时针旋转即可。
// 3.3. p.isRc && x.isRc: g 做一次逆时针旋转即可。
// 3.4. p.isRc && x.isLc: 对 p 做一次顺时针旋转， g 做一次逆时针旋转即可。
func (t *Avl) Insert(key, value interface{}) {
	x := t.searchAt(&t.root, key)
	if *x != nil {
		return
	}
	t.size++
	*x = newBinNode(key, value, t.hot, nil, nil)
	for e := t.hot; e != nil; e = e.parent { // 寻找最低失衡节点 e
		if !e.balanced() {
			t.rotateAt(e.highChild().highChild())
			break
		} else {
			e.updateHeight()
		}
	}
}

// Remove Avl树删除
// 1. 从 _hot 节点出发沿 parent 指针上行，经过O(logn)时间即可确定最低失衡位置 g 位置。
// 2. 作为失衡节点的g ，在不包含 x(被删除节点) 的一侧，必有一个非空孩子 p，且 p 的高度至 少为1。于是，取 g 较高的子节点 p 和孙子节点 v.
// 3. 根据祖孙三代节点 g 、p 和 v 的位置关系，通过以 g 和 p 为轴的适当旋转，同样可以使得这一局部恢复平衡。(同新增)
func (t *Avl) Remove(key interface{}) {
	x := t.searchAt(&t.root, key)
	if *x == nil {
		return
	}
	t.size--
	t.removeAt(x)
	for e := t.hot; e != nil; e = e.parent {
		if !e.balanced() {
			e = t.rotateAt(e.highChild().highChild())
		}
		e.updateHeight()
	}
}
