package tree

// SearchTree 查找树约束接口
type SearchTree interface {
	Size() int                          // 树规模
	Empty() bool                        // 是否空树
	Height() int                        // 树高
	Clear() int                         // 清空查找树并返回清除元素数
	Search(key interface{}) interface{} // 二叉树元素查找
	Insert(key, value interface{})      // 二叉树元素插入
	Remove(key interface{})             // 二叉树元素删除
}

// PrefixTree 前缀树接口
type PrefixTree interface {
	Size() int                                        // 树节点
	Find(key string) interface{}                      // 查找指定节点
	Insert(key string, value interface{}) interface{} // 插入或者替换：返回替换的旧值，插入时返回 nil
	Remove(key string) interface{}                    // 删除节点
	Walk(callable WalkFn)                             // 对所有节点进行处理
	All() []string                                    // 插入或者替换：返回替换的旧值，插入时返回 nil
	ToMap() map[string]interface{}                    // 插入或者替换：返回替换的旧值，插入时返回 nil
	Prefix(prefix string) []string                    // 公共前缀 prefix 的节点列表
	WalkPrefix(prefix string, callable WalkFn)        // 对指定前缀节点进行处理
}
