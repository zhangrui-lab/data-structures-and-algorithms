package tree

// Radix 基数树：在Trie树的基础上新增节点的合并操作即可实现
type Radix struct {
	root *innerNode
	size int
}

// NewRadix 新建空Radix树
func NewRadix() *Radix {
	return &Radix{root: newInnerNode(null)}
}

// 对树节点node进行合并
func (r *Radix) merge(node *innerNode) {
}

// Size 树节点
func (r *Radix) Size() int {
	return r.size
}

// Find 查找指定节点
func (r *Radix) Find(key string) interface{} {
	node := r.find(r.root, key)
	if node == nil || !node.isLeaf() {
		return nil
	}
	return node.leaf.value
}

// 在以 node 为根的子树中查找 key
func (r *Radix) find(node *innerNode, key string) *innerNode {
	for len(key) > 0 && node != nil {
		node = node.getChild(key[0])
		key = key[node.longestCommonPrefix(key):]
	}
	return node
}

// Insert 插入或者替换：返回替换的旧值，插入时返回 nil
func (r *Radix) Insert(key string, value interface{}) interface{} {
	node, search := r.root, key
	for len(search) > 0 {
		child := node.getChild(search[0])
		// 已在树中找到最长公共前缀，剩余部分（search）最为当前key的剩余前缀直接插入
		if child == nil {
			child = node.insertChild(search[0])
			child.prefix = search
		}
		common := child.longestCommonPrefix(search)
		// 存在节点以 search[:common) 为前缀。 此时，需将 search[:common) 分裂为新的公共前缀
		// 1、新建 split 节点，以 child.prefix[common:] 为前缀，以 child.children 为 children。
		// 2、原 child 节点保留 search[:common) 部分作为前缀，以 split 为唯一子节点。
		if common < len(child.prefix) {
			split := newInnerNode(child.prefix[common])
			split.leaf = child.leaf
			split.prefix = child.prefix[common:]
			split.children = child.children

			child.leaf = nil
			child.prefix = child.prefix[:common]
			child.children = nil
			child.children = append(child.children, split)
		}
		search = search[common:]
		node = child
	}
	if node.isLeaf() {
		old := node.leaf.value
		node.leaf.value = value
		return old
	}
	r.size++
	node.insertLeaf(key, value)
	return nil
}

// Remove 删除节点
func (r *Radix) Remove(key string) interface{} {
	value, _ := r.remove(r.root, key) // 需要递归回溯时清空无 key 的子树节点，故采用dfs策略
	return value
}

// 在以node为根的子树中移除 key
func (r *Radix) remove(node *innerNode, key string) (interface{}, bool) {
	if node == nil {
		return nil, false
	}
	if len(key) == 0 {
		if !node.isLeaf() {
			return nil, false
		}
		r.size--
		value := node.leaf.value
		node.leaf = nil
		node.merge()
		return value, node.children.Len() <= 0 && !node.isLeaf() // 不是数据节点且不存在子节点是，当前节点可被移除
	}
	// dfs 深入查询
	child := node.getChild(key[0])
	value, deleted := r.remove(child, key[child.longestCommonPrefix(key):])
	// 空子节点移除
	if deleted {
		node.delChild(key[0])
	}
	// 合并向上传递
	if node != r.root {
		node.merge()
	}
	return value, node.children.Len() <= 0 && !node.isLeaf() // 不是数据节点且不存在子节点是，当前节点可被移除
}

// Walk 对所有节点进行处理
func (r *Radix) Walk(callable WalkFn) {
	r.root.dfs(callable)
}

// All 返回所有节点key
func (r *Radix) All() []string {
	return r.Prefix("")
}

// ToMap 转化为  map<key:value> 结构
func (r *Radix) ToMap() map[string]interface{} {
	var ans = make(map[string]interface{})
	r.Walk(func(key string, value interface{}) bool {
		ans[key] = value
		return false
	})
	return ans
}

// Prefix 公共前缀 prefix 的节点列表
func (r *Radix) Prefix(prefix string) []string {
	var ans []string
	node := r.find(r.root, prefix)
	if node == nil {
		return ans
	}
	node.dfs(func(key string, value interface{}) bool {
		ans = append(ans, key)
		return false
	})
	return ans
}

// WalkPrefix 对指定前缀节点进行处理
func (r *Radix) WalkPrefix(prefix string, callable WalkFn) {
	node := r.find(r.root, prefix)
	if node != nil {
		node.dfs(callable)
	}
}
