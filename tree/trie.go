// Package tree Trie 前缀树
package tree

const null = byte(0x0)

var _ PrefixTree = (*Trie)(nil)

// Trie 前缀树
type Trie struct {
	root *innerNode
	size int
}

// NewTrie 新建空前缀树
func NewTrie() *Trie {
	return &Trie{root: newInnerNode(null)}
}

// Size 树节点
func (t *Trie) Size() int {
	return t.size
}

// Find 查找指定节点
func (t *Trie) Find(key string) (value interface{}) {
	node := t.root.search(key)
	if node != nil && node.isLeaf() {
		return node.leaf.value
	}
	return
}

// Remove 删除节点
func (t *Trie) Remove(key string) interface{} {
	value, _ := t.remove(t.root, key, 0)
	return value
}

// 对于 key=abc, key=abcd 存在且删除 “abc” 时，只会对 “abc” 进行移除
func (t *Trie) remove(node *innerNode, key string, depth int) (interface{}, bool) {
	if node == nil {
		return nil, false
	}
	if len(key) == depth {
		if !node.isLeaf() {
			return nil, false
		}
		t.size--
		value, deleted := node.leaf.value, node.children.Len() <= 0
		node.leaf = nil
		return value, deleted
	}
	value, removed := t.remove(node.getChild(key[depth]), key, depth+1)
	if removed {
		node.delChild(key[depth])
	}
	return value, node.children.Len() <= 0
}

// Insert 插入或者替换：返回替换的旧值，插入时返回 nil
func (t *Trie) Insert(key string, value interface{}) interface{} {
	node := t.root
	for i := 0; i < len(key); i++ {
		parent := node
		node = parent.getChild(key[i])
		if node == nil {
			node = parent.insertChild(key[i])
		}
	}
	if node.isLeaf() {
		old := node.leaf.value
		node.leaf.value = value
		return old
	}
	t.size++
	node.insertLeaf(key, value)
	return nil
}

// Walk 对所有节点进行处理
func (t *Trie) Walk(callable WalkFn) {
	t.root.dfs(callable)
}

// All 返回所有key
func (t *Trie) All() []string {
	return t.Prefix("")
}

// ToMap 将树中节点信息转化为键值对
func (t *Trie) ToMap() map[string]interface{} {
	var ans = make(map[string]interface{})
	t.Walk(func(key string, value interface{}) bool {
		ans[key] = value
		return false
	})
	return ans
}

// Prefix prefix 的节点列表
func (t *Trie) Prefix(prefix string) []string {
	var ans []string
	node := t.root.search(prefix)
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
func (t *Trie) WalkPrefix(prefix string, callable WalkFn) {
	node := t.root.search(prefix)
	if node != nil {
		node.dfs(callable)
	}
}
