package tree

import (
	"data-structures-and-algorithms/contract"
	"data-structures-and-algorithms/queue"
	"fmt"
	"strings"
)

// RedBlack 红黑树（4路平衡搜索树）
// 红黑树约束（规定了红黑树的渐进平衡性）：
// 	(1) 树根始终为黑色
// 	(2) 外部节点均为黑色
// 	(3) 其余节点若为红色，则其孩子节点必为黑色
// 	(4) 从任一外部节点到根节点的沿途，黑节点的数目相等
type RedBlack struct {
	Bst
}

// NewRbTree 新建空RedBlack树
func NewRbTree(cmps ...contract.Comparator) *RedBlack {
	return &RedBlack{Bst: *NewBst(cmps...)}
}

// 双红修正：因新节点 x 的引入，而导致父子节点同为红色的此类情况，称作“双红”（double red）。
// 算法：
// 1.将x的父亲与祖父分别记作p和g。既然此前的红黑树合法，故作为红节点p的父亲，g必然存在且为黑色。g作为内部节点，其另一孩子（即p的兄弟、x的叔父）
//   也必然存在，将其记作u。 以下，视节点u的颜色不同，分两类情况分别处置。
// 2.1 考查u为黑色的情况。此时，x的兄弟、两个孩子的黑高度，均与u相等。从B-树角度等效地看，即同一节点不应包含紧邻
//     的红色关键码，只需令黑色关键码与紧邻的红色关键码互换颜色（伴随着1至2次结构的调整）。如此调整之后，局部子树的黑高度将复原，
//     这意味着全树的平衡也必然得以恢复。 同时，新子树的根节点b为黑色，也不致引发新的双红现象。至此，整个插入操作遂告完成。
// 2.2 节点u为红色的情况。此时，u的左、右孩子非空且均为黑色，其黑高度必与x的兄弟以及两个孩子相等。从B-树角度等效地看，即该节点因超过4度而发生上溢。
//	   只需将B树中g相邻红色节点转为黑色，黑节点 g 转为红色（从B-树的角度来看，等效于上溢节点的一次分裂）。如此调整之后局部子树的黑高度复原。然而，
//	   子树根节点g转为红色之后，有可能在更高层再次引发双红现象。从B-树的角度来看，对应于在关键码g被移出并归入上层节点之后，
//	   进而导致上层节点的上溢，即上溢的向上传播。
// 总结：
//  2.1：2次颜色翻转，2次黑高度更新，1~2次旋转，不在递归。
//  2.2：3次颜色翻转，3次黑高度更新，0次旋转，需要递归。
func (t *RedBlack) solveDoubleRed(x *BinNode) {
	// 根节点，直接变黑并更新高度
	if x.isRoot() {
		x.color = Black
		x.height++
		return
	}
	p := x.parent // 父节点
	// 双红已修正
	if p.isBlack() {
		return
	}
	g, u := p.parent, x.uncle() // 祖父，叔叔(可能为nil)节点
	// 2.1 考查u为黑色的情况
	if u.isBlack() {
		g.color = Red
		if p.isLc() && x.isRc() || p.isRc() && x.isLc() {
			x.color = Black
		} else {
			p.color = Black
		}
		t.rotateAt(x)
		return
	}
	// 2.2 节点u为红色的情况
	p.color = Black
	p.height++
	u.color = Black
	u.height++
	//if !g.isRoot() {
	g.color = Red
	//}
	t.solveDoubleRed(g)
}

// 双黑修正：因某一无红色孩子的黑节点 x 被删除，而导致的此类复杂情况，称作“双黑”（double black） 现象
// 算法：
// 1. 原黑节点x的兄弟必然非空，将其记作s；x的父亲记作p，其颜色不确定。以下视s和p颜色的不同组合，按四种情况分别处置。
// 2.1 节点s为黑且存在红子节点t：则在b树看来，s节点的删除导致下溢，且存在元素足够多的兄弟节点。故，从父节点转借兄弟节点即可。若这三个节点按中序遍
//     历次序重命名为a、b和c，则还需将 a 和c染成黑色，b则继承p此前的颜色。 从红黑树的角度来看， 上述调整过程等效于，对节点t、s和p 实施“3 + 4”重构。
// 2.2.1 节点s及其两个孩子均为黑色时，p为红色时：从B树观察，s节点的删除导致下溢，且兄弟节点也即将下溢，故可从父节点中取出节点g对两个子节点进行合并。
//		 。从红黑树角度看，这一过程可等效地理解为：s和p颜色互换。 经过以上处 理，红黑树的所有条件都在此局部得以恢复。另外，由 于关键码p原为红色，
//		 在关键码p所属节点中，其左或右必然还有一个黑色关键码（当然， 不可能左、右兼有）这意味着，在关键码p从其中取 出之后，不致引发新的下溢。
//		 至此，红黑树条件亦必在全局得以恢复，删除操作即告完成。
// 2.2.2 节点s、s的两个孩子以及节点p均为黑色的情况：在对应的B-树中，因关键码x的删除，导致其所属节点发生下溢。将下溢
//		 节点与其兄弟合并。从红黑树的角度来看，这一过程可等 效地理解为：节点s由黑转红。然而，因s和x在此之前均为黑色，故p原所
//       属的B-树节点必然仅含p这一个关键码。于是在p被借出之后，该节点必将继而发生下溢，故有待于后续 进一步修正。 从红黑树的角度来看，此时的状态可
//       等效地理解为：节点p的父节点刚被删除。
// 2.3 考虑节点s为红色的情况：作为红节点s的父亲，节点p必为黑色；同时，s的两个孩子也应均为黑色。于是从B-树的角度来看，只需，令关键码s与p互换颜色，
//	   即可得到一棵与之完全等价的B-树。而从红黑树的角度来看，这一转换对应于以节点p为轴做一次旋转，并交换节点s与p的颜色。实际上，经过这一转换之后，
//	   情况已经发生了微妙而本质的变化。在转换之后的红黑树中，被删除节点x（及其替代者节点r）有了一个新的兄弟s'。与此前的 兄弟s不同，s'必然是黑的！
//	   这就意味着，接下来可以套用此前所介绍其它情况的处置方法，继续并最终完成双黑修正。
// 总结：
//  2.1：2次颜色翻转，2次黑高度更新，1~2次旋转，不需要递归
//  2.2.1：2次颜色翻转，2次黑高度更新，0次旋转，不需要递归
//  2.2.2：1次颜色翻转，1次黑高度更新，0次旋转，需要递归。
//  2.3：2次颜色翻转，2次黑高度更新，1次旋转，转为2.1或2.21
func (tree *RedBlack) solveDoubleBlack(x *BinNode) {
	var p *BinNode
	if x != nil && x.parent != nil {
		p = x.parent
	}
	if p == nil {
		p = tree.hot
	}
	if p == nil {
		return
	}
	var s *BinNode
	if x == p.lc {
		s = p.rc
	} else {
		s = p.lc
	}
	// 2.3 节点s为红色的情况
	if s.isRed() {
		s.color = Black
		p.color = Red
		if s.isLc() {
			p.rightRotate()
		} else {
			p.leftRotate()
		}
		if s.isRoot() {
			tree.root = s
		}
		tree.solveDoubleBlack(x)
		return
	}
	// 黑兄弟的红子节点
	var t *BinNode
	if s.lc.isRed() {
		t = s.lc
	} else if s.rc.isRed() {
		t = s.rc
	}
	// 2.1：节点s为黑且存在红子节点t
	if t != nil {
		if s.isLc() && t.isRc() || s.isRc() && t.isLc() {
			t.color = p.color
		} else {
			s.color = p.color
			t.color = Black
		}
		p.color = Black
		tree.rotateAt(t)
		return
	}
	// 2.1：节点s为黑且不存在红子节点t
	s.color = Red
	s.height--
	// 2.2.1 节点s及其两个孩子均为黑色时，p为红色时
	if p.isRed() {
		p.color = Black
		return
	}
	// 2.2.2 黑兄弟节点s、s的两个孩子以及节点p均为黑色的情况
	p.height--
	tree.hot = nil
	tree.solveDoubleBlack(p)
	return
}

// Insert 插入 <key, value>，存在时替换
func (t *RedBlack) Insert(key interface{}, value interface{}) {
	x := t.searchAt(&t.root, key)
	if *x != nil {
		(*x).value = value
		return
	}
	t.size++
	*x = newBinNode(key, value, t.hot, nil, nil, Red)
	(*x).height = -1
	t.solveDoubleRed(*x)
}

// Remove 移除存在的key
func (t *RedBlack) Remove(key interface{}) {
	x := t.searchAt(&t.root, key)
	if *x == nil {
		return
	}
	t.size--
	succ := t.removeAt(x)
	if t.hot == nil {
		t.root.color = Black
		t.root.updateHeight()
		return
	}
	if t.balanced(t.hot) {
		return
	}
	if succ.isRed() {
		succ.color = Black
		succ.updateHeight()
		return
	}
	t.solveDoubleBlack(succ)
}

// 当前节点黑高度是否平衡
func (t *RedBlack) balanced(x *BinNode) bool {
	if x.lc.getHeight() != x.rc.getHeight() {
		return false
	}
	if x.isRed() {
		return x.height == x.lc.getHeight()
	}
	return x.height == x.lc.getHeight()+1
}

// 层序打印红黑树基础信息
func (t *RedBlack) levelInfo() string {
	var info []string
	if t.Empty() {
		return ""
	}
	que := queue.New()
	que.Push(t.root)
	for !que.Empty() {
		var tmp string
		size := que.Size()
		for ; size > 0; size-- {
			node := que.Pop().(*BinNode)
			tmp += fmt.Sprintf("{%v,%v,%v,%v}", node.key, node.value, node.color, node.getHeight())
			if node.lc != nil {
				que.Push(node.lc)
			}
			if node.rc != nil {
				que.Push(node.rc)
			}
		}
		tmp = strings.TrimRight(tmp, "\t")
		info = append(info, tmp)
	}
	return strings.Join(info, "\n")
}
