// Package tree 伸展树：基于局部性原理，将被访问的数据亦步亦趋的伸展至根节点，并在伸展过程中对树进行折叠（降低树高，双层伸展）
package tree

import (
	"data-structures-and-algorithms/contract"
)

// Splay 伸展树
type Splay struct {
	Bst
}

// NewSplay 新建空伸展树
func NewSplay(cmps ...contract.Comparator) *Splay {
	return &Splay{Bst: *NewBst(cmps...)}
}

// Search 二叉树元素查找
func (s *Splay) Search(key interface{}) interface{} {
	x := s.searchAt(&s.root, key)
	if *x == nil {
		return nil
	}
	return s.root.value
}

// Insert 二叉树元素插入
// 算法：
// 1. 空树时新建根节点并返回。
// 2. 查找关键码 key， 若存在节点 x，则直接返回。
// 3. 此时最后操作节点 x（hot） 以伸展至树根。x 在中序遍历下为 key 的直接前驱或者直接后继。创建节点 <key, value> y 并根据 x.key 与 y.key 的大小进行操作。
// 4.1. 若 x.key < y.key: 则以 y 为新根，x 为左子节点，x.rc 为右子节点进行更新。
// 4.2. 若 x.key > y.key: 则以 y 为新根，x.lc 为左子节点，x 为右子节点进行更新。
func (s *Splay) Insert(key, value interface{}) {
	if s.Empty() {
		s.size++
		s.root = newBinNode(key, value, nil, nil, nil)
		return
	}
	x := s.searchAt(&s.root, key)
	if *x != nil && (*x).key == key {
		return
	}
	s.size++
	r := s.root
	if keyComparator(r.key, key) < 0 {
		s.root = newBinNode(key, value, nil, r, r.rc)
		if r.rc != nil {
			r.rc.parent = s.root
			r.rc = nil
		}
	} else {
		s.root = newBinNode(key, value, nil, r.lc, r)
		if r.lc != nil {
			r.lc.parent = s.root
			r.lc = nil
		}
	}
	r.parent = s.root
	s.root.updateHeightAbove()
}

// Remove 二叉树元素删除
// 1. 查找关键码 key， 若不存在节点 x，则直接返回。
// 2. 此时，关键码 key 所对应的节点为 s.root。
// 3.1. 若 root 左子为空： 则将 root.rc 设置为根节点即可。
// 3.2. 若 root 右子为空： 则将 root.lc 设置为根节点即可。
// 3.3. 将 root 节点的左右子树分离，并将右子树作为新根，此时在右子树中再次执行关键码 key 的查找，则注定失败，此时会将 key 中序下的直接
// 		后继（右子树皆大于key）伸展至树根。此时，将分离的左子树作为新根的左子树即可。
func (s *Splay) Remove(key interface{}) {
	x := s.searchAt(&s.root, key)
	if *x == nil {
		return
	}
	s.size--
	r := s.root
	if r.lc == nil {
		s.root = r.rc
		r.rc = nil
	} else if r.rc == nil {
		s.root = r.lc
		r.lc = nil
	} else {
		lt := r.lc
		lt.parent = nil
		r.lc = nil
		s.root = r.rc
		s.root.parent = nil
		s.searchAt(&s.root, key)
		s.root.lc = lt
		lt.parent = s.root
	}
	if s.root != nil {
		s.root.parent = nil
		s.root.updateHeight()
	}
}

// searchAt Splay的查找操作，对最后操作的节点（若元素查找成功，则为被查找节点，否则为最后访问节点），需将其双层伸展到根节点处
func (s *Splay) searchAt(x **BinNode, key interface{}) **BinNode {
	r := s.Bst.searchAt(x, key)
	if *r == nil {
		s.splay(s.hot)
	} else {
		s.splay(*r)
	}
	return &s.root
}

// splay 将节点x伸展至根节点
// 双层伸展算法：
// 1. 当前节点为根节点时，终止。
// 2. 当前节点父节点为根节点时，进行一次顺时针旋转（x.IsLc()）或逆时针旋转（x.IsRc()）。
// 3. 其他： 设 p = x.parent, g = p.parent：
// 3.1: p.isLc && x.isLc 时： 对 g 做顺时针旋转， 对 p 做顺时针旋转。
// 3.2: p.isLc && x.isRc 时： 对 p 做逆时针旋转， 对 g 做顺时针旋转。
// 3.3: p.isRc && x.isRc 时： 对 g 做逆时针旋转， 对 p 做逆时针旋转。
// 3.4: p.isRc && x.isLc 时： 对 p 做顺时针旋转， 对 g 做逆时针旋转。
func (s *Splay) splay(x *BinNode) *BinNode {
	if x == nil {
		return nil
	}
	p := x.parent
	var g *BinNode
	if p != nil {
		g = p.parent
	}
	// 情况3
	for p != nil && g != nil {
		if p.isLc() {
			if x.isLc() { // 3.1
				g.rightRotate()
				p.rightRotate()
			} else { // 3.2
				p.leftRotate()
				g.rightRotate()
			}
		} else { // 3.3
			if x.isRc() {
				g.leftRotate()
				p.leftRotate()
			} else { // 3.4
				p.rightRotate()
				x.leftRotate()
			}
		}
		p = x.parent
		if p != nil {
			g = p.parent
		}
	}

	// 情况2
	if x.parent != nil {
		if x.isLc() {
			x.parent.rightRotate()
		} else {
			x.parent.leftRotate()
		}
	}
	x.parent = nil
	s.root = x
	return x
}
