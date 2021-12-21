package tree

// SearchTree 查找树约束接口
type SearchTree interface {
	Size() int                                 // 树规模
	Empty() bool                               // 是否空树
	Height() int                               // 树高
	Clear() int                                // 清空查找树并返回清除元素数
	Search(key, value interface{}) interface{} // 二叉树元素查找
	Insert(key, value interface{})             // 二叉树元素插入
	Remove(key interface{})                    // 二叉树元素删除
}
