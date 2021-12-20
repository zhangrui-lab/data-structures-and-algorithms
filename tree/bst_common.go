package tree

import (
	"data-structures-and-algorithms/types"
	"fmt"
)

// Interface 查找树约束接口
type Interface interface {
	Size() int                                          // 树规模
	Empty() bool                                        // 是否空树
	Height() int                                        // 树高
	Clear() int                                         // 清空查找树并返回清除元素数
	Search(key types.Sortable) (interface{}, error)     // 二叉树元素查找
	Insert(key types.Sortable, value interface{}) error // 二叉树元素插入
	Remove(key types.Sortable) error                    // 二叉树元素删除
}

// Entry 查找树内部词条元素
type Entry struct {
	key   types.Sortable
	value interface{}
}

// Less 比较器支持
func (e Entry) Less(o types.Sortable) bool {
	switch o.(type) {
	case Entry:
		return e.key.Less(o.(Entry).key)
	case types.Sortable:
		return e.key.Less(o)
	default:
		panic(fmt.Sprintf("param need entry or types.Sortable， %t given !", o))
	}
}
