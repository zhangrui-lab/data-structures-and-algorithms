package tree

// Avl Avl树
type Avl struct {
	Bst
}

// NewAvl 新建空avl树
func NewAvl() *Avl {
	return &Avl{}
}

//// Insert Avl树插入
//func (t *Avl) Insert(key types.Sortable, value interface{}) error {
//	x := t.searchAt(&t.root, key)
//	if *x != nil {
//		return fmt.Errorf("key %v already exists", key)
//	}
//	t.size++
//	*x = newBinNode(key, value, t.hot, nil, nil)
//	for e := t.hot; e != nil; e = e.parent {
//		if !e.balanced() {
//			r := t.rotateAt(e.highChild().highChild())
//			*e.fromParent() = r
//			e = r
//			break
//		} else {
//			e.updateHeight()
//		}
//	}
//	return nil
//}
//
//// Remove Avl树删除
//func (t *Avl) Remove(key types.Sortable) error {
//	x := t.searchAt(&t.root, key)
//	if *x == nil {
//		return fmt.Errorf("key %v not found", key)
//	}
//	t.size--
//	t.removeAt(x)
//	for e := t.hot; e != nil; e = e.parent {
//		if !e.balanced() {
//			r := t.rotateAt(e.highChild().highChild())
//			*e.fromParent() = r
//			e = r
//		}
//		e.updateHeight()
//	}
//	return nil
//}
//
//// 在树 t 的合法节点 at（非空较高孙辈节点，其祖父节点产生失衡） 处进行旋转
//func (t *Avl) rotateAt(at *BinNode) *BinNode {
//	var x *BinNode
//	p, g := at.parent, at.parent.parent
//	if p.isLc() {
//		if at.isLc() {
//			p.parent = g.parent
//			x = connect34(at, p, g, at.lc, at.rc, p.rc, g.rc)
//		} else {
//			at.parent = g.parent
//			x = connect34(p, at, g, p.lc, at.lc, at.rc, g.rc)
//		}
//	} else {
//		if at.isRc() {
//			p.parent = g.parent
//			x = connect34(g, p, at, g.lc, p.lc, at.lc, at.rc)
//		} else {
//			at.parent = g.parent
//			x = connect34(g, at, p, g.lc, at.lc, at.rc, p.rc)
//		}
//	}
//	return x
//}
