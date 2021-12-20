// Package tree 通用二叉树节点
package tree

import (
	"data-structures-and-algorithms/queue"
	"data-structures-and-algorithms/stack"
	"data-structures-and-algorithms/types"
)

type RbColor int

const (
	_ RbColor = iota
	Red
	Black
)

// BinNode 二叉树节点
// 当前节点为二叉树通用节点：会维护各种树结构的冗余信息
// 约定节点深度为从根节点到当前节点所过边的长度。 节点高度为以其为根的子树的最大深度。
// 约定：根节点高度为0， 外部节点（nil节点）高度为 -1
type BinNode struct {
	data           types.Sortable
	parent, lc, rc *BinNode // 通用数据信息
	height         int      // avl树中用以计算平衡因子；左式堆中用以保存 Null Path Length。
	color          RbColor  // 红黑树颜色：默认以红节点给出
}

// newBinNode  新建普通二叉树节点
func newBinNode(data types.Sortable) *BinNode {
	return &BinNode{data: data, height: 0, color: Red}
}

// newBstNode 以<Key, Value> 形式新建二叉查找树节点
func newBstNode(key types.Sortable, value interface{}) *BinNode {
	return &BinNode{data: &entry{key: key, value: value}, height: 0, color: Red}
}

// isRoot 是否可为根节点。约定：e != nil
func (e *BinNode) isRoot() bool {
	return e.parent == nil
}

// isLc 是否作为父节点的左子节点。约定：e != nil
func (e *BinNode) isLc() bool {
	return !e.isRoot() && e.parent.lc == e
}

// isRc 是否作为父节点的右子节点。约定：e != nil
func (e *BinNode) isRc() bool {
	return !e.isRoot() && e.parent.rc == e
}

// isLc 至少拥有一个孩子。约定：e != nil
func (e *BinNode) hasChild() bool {
	return e.lc != nil || e.rc != nil
}

// hasBothChild 同时拥有两个孩子。约定：e != nil
func (e *BinNode) hasBothChild() bool {
	return e.lc != nil && e.rc != nil
}

// isLeaf 当前节点是否为叶节点
func (e *BinNode) isLeaf() bool {
	return !e.hasChild()
}

// sibling 当前节点的兄弟节点。约定：e != nil && e.parent != nil
func (e *BinNode) sibling() *BinNode {
	if e.isLc() {
		return e.parent.rc
	}
	return e.parent.lc
}

// uncle 当前节点的叔叔节点（父节点的兄弟节点）。约定：e != nil && e.parent != nil && e.parent.parent != nil
func (e *BinNode) uncle() *BinNode {
	if e.parent.isLc() {
		return e.parent.parent.rc
	}
	return e.parent.parent.lc
}

// fromParentTo 返回来自e父节点的指针信息。用以重设 e 父节点到当前节点的指针。约定：e != nil && e.parent != nil
func (e *BinNode) fromParentTo() **BinNode {
	if e.isLc() {
		return &(e.parent.lc)
	}
	return &(e.parent.rc)
}

// insertAsLc 将 v 作为 e 的左子节点插入 （e左子节点不存在）
func (e *BinNode) insertAsLc(v types.Sortable) *BinNode {
	node := &BinNode{
		data:   v,
		parent: e,
	}
	e.lc = node
	return node
}

// insertAsRc 将 v 作为 e 的右子节点插入 （e右子节点不存在）
func (e *BinNode) insertAsRc(v types.Sortable) *BinNode {
	node := &BinNode{
		data:   v,
		parent: e,
	}
	e.rc = node
	return node
}

// size 以 e 为根的子树的元素个数
func (e *BinNode) size() int {
	if e == nil {
		return 0
	}
	return 1 + e.lc.size() + e.rc.size()
}

// getHeight 获取当前节点高度
func (e *BinNode) getHeight() int {
	if e == nil {
		return -1
	}
	return e.height
}

// succ 获取当前节点中序遍历下的直接后继：后继存在时返回该节点，不存在时（最右侧节点）返回nil
// 算法描述：1、右子树存在时，位于右子树的最左侧节点中；2、右子树不存在时，位于将当前节点最为左子树节点的最低节点中。
func (e *BinNode) succ() *BinNode {
	if e.rc != nil {
		for e = e.rc; e.lc != nil; e = e.lc {
		}
	} else {
		for ; e != nil && e.isRc(); e = e.parent {
		}
		if e != nil {
			e = e.parent
		}
	}
	return e
}

// balanced 是否平衡：节点的左子树的高度与右子树的高度差不超过 1.
func (e *BinNode) balanced() bool {
	if e == nil {
		return true
	}
	diff := e.balanceFac()
	return -2 < diff && diff < 2
}

// balanceFac 平衡因子
func (e *BinNode) balanceFac() int {
	return e.lc.getHeight() - e.rc.getHeight()
}

// updateHeight 对当前非空节点的高度进行更新
func (e *BinNode) updateHeight() {
	h := e.lc.getHeight()
	rh := e.rc.getHeight()
	if h < rh {
		h = rh
	}
	e.height = h + 1
}

// updateHeightAbove 更新高度, 从x出发，覆盖历代祖先。
func (e *BinNode) updateHeightAbove() {
	for e != nil {
		h := e.getHeight()
		e.updateHeight()
		if h == e.getHeight() {
			break
		}
		e = e.parent
	}
}

// rightRotate 对节点右旋（顺时针）：成功的右旋会令其合法左子节点接替当前节点位置，当前节点成为其左子节点的右子节点；
func (e *BinNode) rightRotate() *BinNode {
	if e == nil || e.lc == nil {
		return nil
	}
	lc := e.lc
	lc.parent = e.parent
	if e.parent != nil {
		*e.fromParentTo() = lc
	}
	e.parent = lc

	e.lc = lc.rc
	if e.lc != nil {
		e.lc.parent = e
	}
	lc.rc = e
	return nil
}

// leftRotate 对空节点左旋（逆时针）：成功的左旋会令其合法右子节点接替当前节点位置，当前节点成为其右子节点的左子节点；
func (e *BinNode) leftRotate() *BinNode {
	if e == nil || e.rc == nil {
		return nil
	}
	rc := e.rc
	rc.parent = e.parent
	if e.parent != nil {
		*e.fromParentTo() = rc
	}
	e.parent = rc
	e.rc = rc.lc
	if e.rc != nil {
		e.rc.parent = e
	}
	rc.lc = e
	return interface{}(rc).(*BinNode)
}

// travelLevel 对以当前节点为根的子树进行层序遍历
func (e *BinNode) travelLevel(visit func(sortable *types.Sortable)) {
	que := queue.New()
	que.Push(e)
	for !que.Empty() {
		e, _ = que.Pop().(*BinNode)
		visit(&e.data)
		if e.lc != nil {
			que.Push(e.lc)
		}
		if e.rc != nil {
			que.Push(e.rc)
		}
	}
}

// travelPre 对以当前节点为根的子树进行先序遍历
func (e *BinNode) travelPre(visit func(sortable *types.Sortable)) {
	e.stackPre1(visit)
}

// dfsPre 递归版先序遍历
func (e *BinNode) dfsPre(visit func(sortable *types.Sortable)) {
	if e == nil {
		return
	}
	visit(&e.data)
	e.lc.dfsPre(visit)
	e.rc.dfsPre(visit)
}

// stackPre1 Stack迭代版1 先序遍历
func (e *BinNode) stackPre1(visit func(sortable *types.Sortable)) {
	stk := stack.New()
	goLeftAndVisit := func(x *BinNode) {
		for ; x != nil; x = x.lc {
			visit(&x.data)
			if x.rc != nil {
				stk.Push(x.rc)
			}
		}
	}
	for {
		goLeftAndVisit(e)
		if stk.Empty() {
			break
		}
		e = stk.Pop().(*BinNode)
	}
}

// stackPre2 Stack迭代版2 先序遍历：考虑给定节点和其左右子节点访问顺序为 r > r.lc > r.rc， 故逆序入栈即可
func (e *BinNode) stackPre2(visit func(sortable *types.Sortable)) {
	stk := stack.New()
	stk.Push(e)
	for !stk.Empty() {
		e = stk.Pop().(*BinNode)
		visit(&e.data)
		if e.rc != nil {
			stk.Push(e.rc)
		}
		if e.lc != nil {
			stk.Push(e.lc)
		}
	}
}

// travelPre 对以当前节点为根的子树进行中序遍历
func (e *BinNode) travelIn(visit func(sortable *types.Sortable)) {
	e.stackIn2(visit)
}

// dfsIn 递归版中序遍历
func (e *BinNode) dfsIn(visit func(sortable *types.Sortable)) {
	if e == nil {
		return
	}
	e.lc.dfsIn(visit)
	visit(&e.data)
	e.rc.dfsIn(visit)
}

// stackIn1 栈迭代版中序1
// 算法描述：从当前节点出发，沿左分支不断深入并入栈，直至没有左分支的节点。随后弹出栈顶节点访问之并转向右子树。
func (e *BinNode) stackIn1(visit func(sortable *types.Sortable)) {
	stk := stack.New()
	goLeft := func(x *BinNode) {
		for ; x != nil; x = x.lc {
			stk.Push(x)
		}
	}
	for {
		goLeft(e)
		if stk.Empty() {
			break
		}
		e = stk.Pop().(*BinNode)
		visit(&e.data)
		e = e.rc
	}
}

// stackIn2 栈迭代版中序2
// 算法描述（可参考迭代1）：从根节点开始深入遍历左子树并入栈。当无左子节点时，从栈中弹出元素并方位，若栈为空，则代表遍历完成。
func (e *BinNode) stackIn2(visit func(sortable *types.Sortable)) {
	stk := stack.New()
	for {
		if e != nil {
			stk.Push(e)
			e = e.lc
		} else if !stk.Empty() {
			e = stk.Pop().(*BinNode)
			visit(&e.data)
			e = e.rc
		} else {
			break
		}
	}
}

// backtrackIn 回溯中序迭代版
// 算法描述：
//	1. 不存在回溯标记且存在左子树时，深入至最左侧节点 x，访问该节点 x。
//	2. 若 x 存在右子树，转向右子树，清除回溯标志，并执行步骤 1。
//	3. 若 x 不存在右子树，尝试对其进行回溯（回溯回中序遍历序列下的直接后继，在树拓扑结构中位于将其作为左子树的最低节点），并设置回溯标记。
func (e *BinNode) backtrackIn(visit func(sortable *types.Sortable)) {
	back := false
	for e != nil {
		if !back && e.lc != nil {
			e = e.lc
		} else {
			visit(&e.data)
			if e.rc != nil {
				e = e.rc
				back = false
			} else {
				e = e.succ()
				back = true
			}
		}
	}
}

// iterIn 迭代版中序：无需stack和回溯标记
// 算法描述：
// 1、从根节点处深入最左侧节点 x，访问节点x。
// 2、转向后续访问对象（x的后继：存在与x的右子树最左侧节点或祖先节点中）：
// 	2.1、x 存在右子树，x=x.rc，并执行步骤 1。
// 	2.1、x 不存在右子树，转向后继节点，并访问之，若后继节点 p 存在右子树，停止回溯并令 x=p.rc，并执行步骤 1。
// 3、x 为 nil 时终止（从树的最右侧节点回溯到根节点的父节点：nil）。
func (e *BinNode) iterationIn(visit func(sortable *types.Sortable)) {
	for e != nil {
		if e.lc != nil {
			e = e.lc
			continue
		}
		visit(&e.data)
		if e.rc != nil {
			e = e.rc
			continue
		}
		for e = e.succ(); e != nil && e.rc == nil; e = e.succ() {
			visit(&e.data)
		}
		if e != nil {
			visit(&e.data)
			e = e.rc
		}
	}
}

// travelPost 对以当前节点为根的子树进行后序遍历
func (e *BinNode) travelPost(visit func(sortable *types.Sortable)) {
	e.morrisPost(visit)
}

// dfsPost 递归版后续遍历
func (e *BinNode) dfsPost(visit func(sortable *types.Sortable)) {
	if e == nil {
		return
	}
	e.lc.dfsPost(visit)
	e.rc.dfsPost(visit)
	visit(&e.data)
}

// stackPost 栈迭代版后序
// 算法：
// 1、从根节点向最左侧路径进行深入并将沿途几点入栈：入栈优先级为 当前节点 > 存在的右子节点 > 存在的右子节点
// 2、栈不为空时，移除栈顶元素 x 并访问。
//	2.1、若此时栈中还有元素，考虑入栈顺序，栈顶元素必为 x 的右兄弟节点或者父节点。若栈顶为右兄弟节点时， 以栈顶元素启动步骤1；若为父节点时，不做操作；
//	2.1、若此时栈中无元素，不做操作。
func (e *BinNode) stackPost(visit func(sortable *types.Sortable)) {
	stk := stack.New()
	goLeft := func(x *BinNode) {
		for ; x != nil; x = x.lc {
			stk.Push(x)
			if x.rc != nil {
				stk.Push(x.rc)
			}
		}
	}
	goLeft(e)
	for !stk.Empty() {
		e = stk.Pop().(*BinNode)
		visit(&e.data)
		if !stk.Empty() && e.parent != stk.Top().(*BinNode) {
			goLeft(stk.Pop().(*BinNode))
		}
	}
}

// Morris 遍历算法
// Morris 遍历算法是另一种遍历二叉树的方法，它能将非递归的中序遍历空间复杂度降为 O(1)。

// morrisIn 中序遍历
// 算法整体步骤如下（假设当前遍历到的节点为 x）：
// 1、如果 x 无左孩子，访问 x，再访问 x 的右孩子，即 x = x.rc。
// 2、如果 x 有左孩子，则找到 x 左子树上最右的节点（即左子树中序遍历的最后一个节点，x 在中序遍历中的前驱节点），
//	  我们记为 predecessor。根据 predecessor 的右孩子是否为空，进行如下操作。
// 	 2.1、如果 predecessor 的右孩子为空，则将其右孩子指向 x，然后访问 x 的左孩子，即 x = x.lc。
// 	 2.2、如果 predecessor 的右孩子不为空，则此时其右孩子指向 x，说明我们已经遍历完 x 的左子树，我们将
//	 	  predecessor 的右孩子置空，将 x 的值加入答案数组，然后访问 x 的右孩子，即 x = x.rc。
func (e *BinNode) morrisIn(visit func(sortable *types.Sortable)) {
	for e != nil {
		if e.lc == nil { // 无左子，访问当前节点并深入右子
			visit(&e.data)
			e = e.rc
			continue
		}
		predecessor := lTreePred(e)
		if predecessor.rc == nil { // 左子尚未开始访问
			predecessor.rc = e
			e = e.lc
			continue
		}
		visit(&e.data) // 左子访问完成，可访问当前节点
		predecessor.rc = nil
		e = e.rc
	}
}

// morrisPre 前序遍历
func (e *BinNode) morrisPre(visit func(sortable *types.Sortable)) {
	for e != nil {
		if e.lc == nil {
			visit(&e.data)
			e = e.rc
			continue
		}
		predecessor := lTreePred(e)
		if predecessor.rc == nil {
			visit(&e.data)
			predecessor.rc = e
			e = e.lc
			continue
		}
		predecessor.rc = nil
		e = e.rc
	}
}

// morrisPre 后序遍历
func (e *BinNode) morrisPost(visit func(sortable *types.Sortable)) {
	reverseVisit := func(p *BinNode, x *BinNode) {
		for ; p != x; p = p.parent {
			visit(&p.data)
		}
		visit(&p.data)
	}
	r := e
	for {
		if e.lc != nil {
			predecessor := lTreePred(e)
			if predecessor.rc == nil {
				predecessor.rc = e
				e = e.lc
				continue
			} else {
				predecessor.rc = nil
				reverseVisit(predecessor, e.lc)
			}
		}
		if e.rc == nil {
			break
		}
		e = e.rc
	}
	reverseVisit(e, r)
}

// lTreePred 在左子树中寻找节点x中序下的直接前驱节点，仅供morris算法调用
func lTreePred(x *BinNode) *BinNode {
	predecessor := x.lc
	for ; predecessor.rc != nil && predecessor.rc != x; predecessor = predecessor.rc {
	}
	return predecessor
}
