// Package tree 词条类型支持：k-v数据类型的内置保存格式。
package tree

import (
	"data-structures-and-algorithms/types"
	"fmt"
)

// entry 词条类型
type entry struct {
	key   types.Sortable
	value interface{}
}

// 比较器支持
func (e *entry) Less(o types.Sortable) bool {
	switch o.(type) {
	case *entry:
		return e.key.Less(o.(*entry).key)
	case types.Sortable:
		return e.key.Less(o)
	default:
		panic(fmt.Sprintf("param need entry or types.Sortable， %t given !", o))
	}
}
