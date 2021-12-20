package tree

import (
	"data-structures-and-algorithms/types"
	"fmt"
)

// btreeNode b树节点
type btreeNode struct {
	parent *btreeNode
	key    []types.Sortable
	child  []*btreeNode
}

// Btree B-树
type Btree struct {
	size  int
	order int
	hot   *btreeNode
	root  *btreeNode
}

func (t *Btree) Size() int {
	return t.size
}

func (t *Btree) Height() int {
	return 0
}

func (t *Btree) Empty() bool {
	return t.size <= 0
}

func (t *Btree) Order() int {
	return t.order
}

// 清空查找树并返回清除元素数
func (t *Btree) Clear() int {
	size := t.size
	size = 0
	t.root = nil
	return size
}

// 二叉树元素查找
func (t *Btree) Search(key types.Sortable) (interface{}, error) {
	if t.Empty() {
		return nil, fmt.Errorf("search %v in empty btree", key)
	}
	x := t.root
	for x != nil {
		//v, _ := x.key.At(p)
		//if v == key {
		//	return v, nil
		//}
		//r, _ = x.child.At(p + 1)
		//x = r.
	}
	return nil, fmt.Errorf("key %v not found", key)
}

// 二叉树元素插入
//func (t *Btree) Insert(key types.Sortable, value interface{}) error {
//
//}
//
//// 二叉树元素删除
//func (t *Btree) Remove(key types.Sortable) error {
//
//}
//
//// 因插入而上溢之后的分裂处理
////
//func (t *Btree) solveOverflow(x *btreeNode) {
//}
//
//// 因删除而下溢之后的合并处理
//func (t *Btree) solveUnderflow(x *btreeNode) {
//}
