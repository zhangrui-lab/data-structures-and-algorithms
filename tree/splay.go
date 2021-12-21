package tree

// Splay 伸展树
type Splay struct {
	Bst
}

//// Search 二叉树元素查找
//func (s *Splay) Search(key types.Sortable) (interface{}, error) {
//	x := s.searchAt(&s.root, key)
//	if *x == nil {
//		return nil, fmt.Errorf("key %v not found", key)
//	}
//	return s.root.data.(Entry).value, nil
//}
//
//// Insert 二叉树元素插入
//func (s *Splay) Insert(key types.Sortable, value interface{}) error {
//	if s.Empty() {
//		s.size++
//		s.root = newBstNode(key, value, nil, nil, nil)
//		return nil
//	}
//	x := s.searchAt(&s.root, key)
//	if *x != nil && (*x).data.(Entry).key == key {
//		return fmt.Errorf("key %v already exists", key)
//	}
//	s.size++
//	r := s.root
//	if r.data.(Entry).key.Less(key) {
//		r.parent = newBstNode(key, value, nil, r, r.rc)
//		s.root = r.parent
//		if r.rc != nil {
//			r.rc.parent = s.root
//			r.rc = nil
//		}
//	} else {
//		r.parent = newBstNode(key, value, nil, r.lc, r)
//		s.root = r.parent
//		if r.lc != nil {
//			r.lc.parent = s.root
//			r.lc = nil
//		}
//	}
//	r.updateHeightAbove()
//	return nil
//}
//
//// Remove 二叉树元素删除
//func (s *Splay) Remove(key types.Sortable) error {
//	if s.Empty() {
//		return fmt.Errorf("empty splay tree")
//	}
//	x := s.searchAt(&s.root, key)
//	if *x == nil {
//		return fmt.Errorf("key %v not found", key)
//	}
//	r := s.root
//	if r.lc == nil {
//		s.root = r.rc
//		r.rc = nil
//	} else if r.rc == nil {
//		s.root = r.lc
//		r.lc = nil
//	} else {
//		lt := r.lc
//		lt.parent = nil
//		r.lc = nil
//		s.root = r.rc
//		s.root.parent = nil
//		s.searchAt(&s.root, key)
//		s.root.lc = lt
//		lt.parent = s.root
//	}
//	s.size--
//	if s.root != nil {
//		s.root.parent = nil
//		s.root.updateHeight()
//	}
//	return nil
//}
//
//// searchAt Splay的查找操作，对最后操作的节点，需将其双层伸展到根节点处
//func (s *Splay) searchAt(x **BinNode, v types.Sortable) **BinNode {
//	r := s.Bst.searchAt(x, v)
//	if *r == nil {
//		s.splay(s.hot)
//	}
//	s.splay(*r)
//	return r
//}

// splay 将节点x伸展至根节点
// 双层伸展算法：
// 1. 当前节点为根节点时，终止。
// 2. 当前节点父节点为根节点时，进行一次顺时针旋转（x.IsLc()）或逆时针旋转（x.IsRc()）。
// 3. 其他： 设 p = x.parent, g = p.parent
// 3.1: x.isLc && p.isLc 时
func (s *Splay) splay(x *BinNode) *BinNode {
	if x == nil {
		return nil
	}
	p := x.parent
	var g *BinNode
	if p != nil {
		g = p.parent
	}
	for p != nil && g != nil {
		if p.isLc() {
			if x.isLc() {
				g.rightRotate()
				p.rightRotate()
			} else {
				p.leftRotate()
				g.rightRotate()
			}
		} else {
			if x.isRc() {
				g.leftRotate()
				p.leftRotate()
			} else {
				p.rightRotate()
				x.leftRotate()
			}
		}
		p = x.parent
		if p != nil {
			g = p.parent
		}
	}

	if x.parent != nil {
		if x.isLc() {
			x.parent.rightRotate()
		} else {
			x.parent.leftRotate()
		}
	}
	x.parent = nil
	return x
}
