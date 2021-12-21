// Package tree 二叉搜索树：在二叉树的基础上新增节点局部约束即可实现，任意节点x，x.lc.data <= x.data <= x.rc.data
package tree

import (
	"data-structures-and-algorithms/contract"
	"fmt"
	"strings"
)

var (
	keyComparator = contract.DefaultComparator
)

// Bst 二叉查找树：不维护树的渐进平衡性
type Bst struct {
	hot           *BinNode            // “命中”节点的父亲
	root          *BinNode            // 根节点
	keyComparator contract.Comparator // 比较器
	size          int
}

// NewBst 创建二叉搜索树
func NewBst(cmps ...contract.Comparator) *Bst {
	cmp := keyComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	return &Bst{keyComparator: cmp}
}

// Size 树规模
func (t *Bst) Size() int {
	return t.size
}

// Empty 是否为空树
func (t *Bst) Empty() bool {
	return t.size <= 0
}

// Height 树高
func (t *Bst) Height() int {
	return t.root.getHeight()
}

// Clear 清空二叉查找树
func (t *Bst) Clear() int {
	size := t.size
	t.hot = nil
	t.root = nil
	t.size = 0
	return size
}

// Search 二叉树元素查找
func (t *Bst) Search(key interface{}) interface{} {
	x := t.searchAt(&t.root, key)
	if *x == nil {
		return nil
	}
	return (*x).value
}

// Insert 二叉树元素插入
func (t *Bst) Insert(key interface{}, value interface{}) {
	x := t.searchAt(&t.root, key)
	if *x != nil {
		return
	}
	t.size++
	*x = newBinNode(key, value, t.hot, nil, nil)
	t.hot.updateHeightAbove()
}

// Remove 二叉树元素删除
func (t *Bst) Remove(key interface{}) {
	x := t.searchAt(&t.root, key)
	if *x == nil {
		return
	}
	t.size--
	t.removeAt(x)
	t.hot.updateHeightAbove()
}

// searchAt 在以x为根节点的子树中查找元素v，设置hot指针, 并返回元素所在位置指针（指针的指针，便于上层直接赋值）
func (t *Bst) searchAt(x **BinNode, key interface{}) **BinNode {
	t.hot = nil
	if *x != nil {
		t.hot = (*x).parent
	}
	for !equal(*x, key) {
		t.hot = *x
		if keyComparator(key, (*x).key) < 0 {
			x = &(*x).lc
		} else {
			x = &(*x).rc
		}
	}
	return x
}

// removeAt 从树t中摘除节点x：返回值指向实际被删除节点的接替者（中序下的直接后继）；hot指向实际被删除节点的父亲
// 算法描述：
// 1、 若当前节点无左子，则将当前节点替换为其右子。（无子结点也被囊括）
// 2、 若当前节点无右子，则将当前节点替换为其左子。
// 3、 双子俱全时：
// 	3.1、 在其右子树中定位其直接后继元素 p。
// 	3.2、 交换 x 与 p。删除 p。
func (t *Bst) removeAt(x **BinNode) *BinNode {
	w := *x           // 实际被删除节点
	var succ *BinNode // 后继节点
	if (*x).lc == nil {
		*x = (*x).rc
		succ = *x
	} else if (*x).rc == nil {
		*x = (*x).lc
		succ = *x
	} else {
		w = (*x).successor()
		// todo 此处未作节点交换，只实现对数据项的交换 (外层节点引用的数据信息会出现异常)
		(*x).key, (*x).value = w.key, w.value
		p := w.parent
		if p == (*x) {
			p.rc = w.rc
		} else {
			p.lc = w.rc
		}
		succ = w.rc
	}
	t.hot = w.parent
	if succ != nil { // 设置子节点指针
		succ.parent = t.hot
	}
	return succ
}

// String 中序遍历下输出树元素
func (t *Bst) String() string {
	items := make([]string, 0, t.Size())
	t.root.travelIn(func(key, value interface{}) {
		items = append(items, fmt.Sprintf("{%v,%v}", key, value))
	})
	return "{" + strings.Join(items, ", ") + "}"
}

// equal 节点判等：外部节点假想为通配符哨兵
func equal(x *BinNode, key interface{}) bool {
	return x == nil || x.key == key
}
