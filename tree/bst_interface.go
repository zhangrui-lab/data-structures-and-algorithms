package tree

import "data-structures-and-algorithms/types"

// Interface 查找树约束接口
type Interface interface {
	Clear()                           // 清空查找树
	Search(v types.Sortable) *BinNode // 二叉树元素查找
	Insert(v types.Sortable) *BinNode // 二叉树元素插入
	Remove(v types.Sortable) bool     // 二叉树元素删除
}
