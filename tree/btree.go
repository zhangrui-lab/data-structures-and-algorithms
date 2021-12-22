package tree

import (
	"data-structures-and-algorithms/contract"
	"data-structures-and-algorithms/queue"
	"fmt"
	"sort"
	"strings"
)

var (
	nilItems    = make([]interface{}, 16)
	nilChildren = make([]*btreeNode, 16)
)

// b-树节点
//   k0  k1  k2
// c0  c1  c2  c3
// 1. len(key) = len(children) - 1
// 2. ci <= ki < ci+1
// 3. 空 btreeNode 无数据信息，包含一个为 nil 的 lc。
type btreeNode struct {
	parent   *btreeNode
	items    []interface{}
	children []*btreeNode
}

// 新建空节点
func emptyBtreeNode(parent *btreeNode) *btreeNode {
	node := &btreeNode{parent: parent}
	node.children = append(node.children, nil)
	return node
}

// 在指定位置插入元素
func (n *btreeNode) insertItemAt(index int, item interface{}) {
	n.items = append(n.items, nil)
	if index < len(n.items) {
		copy(n.items[index+1:], n.items[index:])
	}
	n.items[index] = item
}

// 在指定位置移除元素
func (n *btreeNode) removeItemAt(index int) interface{} {
	item := n.items[index]
	copy(n.items[index:], n.items[index+1:])
	n.items[len(n.items)-1] = nil
	n.items = n.items[:len(n.items)-1]
	return item
}

// 在指定位置插入子节点指针
func (n *btreeNode) insertChildAt(index int, child *btreeNode) {
	n.children = append(n.children, nil)
	if index < len(n.children) {
		copy(n.children[index+1:], n.children[index:])
	}
	n.children[index] = child
}

// 在指定位置删除子节点指针
func (n *btreeNode) removeChildAt(index int) *btreeNode {
	child := n.children[index]
	copy(n.children[index:], n.children[index+1:])
	n.children[len(n.children)-1] = nil
	n.children = n.children[:len(n.children)-1]
	return child
}

// 截断slice中的元素
func (n *btreeNode) truncateItem(index int) {
	var toClear []interface{}
	n.items, toClear = n.items[:index], n.items[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilItems):]
	}
}

// 截断slice中子节点指针元素
func (n *btreeNode) truncateChild(index int) {
	var toClear []*btreeNode
	n.children, toClear = n.children[:index], n.children[index:]
	for len(toClear) > 0 {
		toClear = toClear[copy(toClear, nilChildren):]
	}
}

// 节点 n 在位置 index 处进行分裂
func (n *btreeNode) split(index int) (interface{}, *btreeNode) {
	item := n.items[index]
	node := &btreeNode{}
	node.items = append(node.items, n.items[index+1:]...)
	node.children = append(node.children, n.children[index+1:]...)
	n.truncateItem(index)
	n.truncateChild(index + 1)
	if node.children[0] != nil {
		for _, child := range node.children {
			child.parent = node
		}
	}
	return item, node
}

// 当前节点为其父节点的第几个孩子
func (n *btreeNode) nth() int {
	var index int
	for size := len(n.parent.children); index < size; index++ {
		if n.parent.children[index] == n {
			break
		}
	}
	return index
}

// 将节点 n 与其右兄弟以及父节点中 index 位置的元素进进行合并.
func (n *btreeNode) merge(index int, rs *btreeNode) *btreeNode {
	item := n.parent.removeItemAt(index)
	n.parent.removeChildAt(index + 1)
	n.insertItemAt(len(n.items), item)
	n.items = append(n.items, rs.items...)
	n.children = append(n.children, rs.children...)
	rs.items = nil
	rs.children = nil
	return n
}

// 在树节点中寻找 item 所在位置：成功时返回 item[i] == key； 失败时 item[i-1] < key < item[i]
func (n *btreeNode) find(comparator contract.Comparator, item interface{}) (int, bool) {
	i := sort.Search(len(n.items), func(i int) bool {
		return comparator(item, n.items[i]) < 0
	})
	if i > 0 && comparator(item, n.items[i-1]) == 0 {
		return i - 1, true
	}
	return i, false
}

// Btree B-树
type Btree struct {
	hot           *btreeNode
	root          *btreeNode
	size          int
	order         int
	keyComparator contract.Comparator
}

// 算法：
// 1、x != nil 时，在 x 中查找 key 对于位置 i。
// 2.1、x.key[i] == key， return x。
// 2.2、x.key[i] != key， x = x.children[i]，执行 1。
func (t *Btree) searchAt(node *btreeNode, item interface{}) (*btreeNode, int) {
	t.hot = nil
	for node != nil {
		i, found := node.find(t.keyComparator, item)
		if found {
			return node, i
		}
		t.hot = node
		node = node.children[i]
	}
	return nil, 0
}

// solveOverflow 上溢修复
// 算法：
// 1、节点 x 是否上溢，未上溢则终止。
// 2、节点元素取中 item, 并分裂为两个节点 l, r。处理 l， r 中的 keys 和 children 以及子节点指针。
// 4、在x父节点(父节点不存在时则意味着在根节点溢出，新建根节点)适当位置插入 item，以及新的子节点指针并指向 r。
// 4、令 x 为 x父节点，执行算法 1。
func (t *Btree) solveOverflow(node *btreeNode) {
	if len(node.children) <= t.order {
		return
	}
	index := t.order >> 1
	item, newNode := node.split(index)
	p := node.parent
	if p == nil {
		p = emptyBtreeNode(nil)
		p.children[0] = node
		node.parent = p
		t.root = p
	}
	newNode.parent = p
	r, _ := p.find(t.keyComparator, item)
	p.insertItemAt(r, item)
	p.insertChildAt(r+1, newNode)
	t.solveOverflow(p)
}

// solveUnderflow 下溢修复
// 算法（左顾右盼+合并）：
// 1、节点 x 是否下溢，未下溢则终止。
// 2、x 为树根，倘若作为树根的节点已不含关键码，却有（唯一的）非空孩子，则这个节点可被跳过,并因不再有用而被销毁,整树高度降低一层。
// 3、x左兄弟存在且元素充足，从左兄弟转借一个元素。
// 4、x右兄弟存在且元素充足，从右兄弟转借一个元素。
// 5、合并x与左或右兄弟。
func (t *Btree) solveUnderflow(node *btreeNode) {
	if (t.order+1)>>1 <= len(node.children) {
		return
	}
	p := node.parent
	if p == nil {
		if len(node.items) == 0 && node.children[0] != nil {
			t.root = node.children[0]
			t.root.parent = nil
			node.children[0] = nil
		}
		return
	}
	index := node.nth()
	if index < len(p.children)-1 && (t.order+1)>>1 < len(p.children[index+1].children) { // 从右孩子借
		rs := p.children[index+1]
		node.insertItemAt(len(node.items), p.items[index])
		p.items[index] = rs.removeItemAt(0)
		child := rs.removeChildAt(0)
		node.insertChildAt(len(node.children), child)
		if child != nil {
			child.parent = node
		}
		return
	} else if index > 0 && (t.order+1)>>1 < len(p.children[index-1].children) { // 从左孩子借
		ls := p.children[index-1]
		node.insertItemAt(0, p.items[index-1])
		p.items[index-1] = ls.removeItemAt(len(ls.items) - 1)
		child := ls.removeChildAt(len(ls.children) - 1)
		node.insertChildAt(0, child)
		if child != nil {
			child.parent = node
		}
		return
	} else if index < len(p.children)-1 { // 与右孩子合并
		node.merge(index, p.children[index+1])
	} else { // 与左孩子合并
		p.children[index-1].merge(index-1, node)
	}
	t.solveUnderflow(p)
}

// NewBtree 初始化空B-树
func NewBtree(order int, cmps ...contract.Comparator) *Btree {
	cmp := contract.DefaultComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	return &Btree{order: order, keyComparator: cmp}
}

// Size 元素个数
func (t *Btree) Size() int {
	return t.size
}

// Height 树高
func (t *Btree) Height() (height int) {
	if t.root == nil {
		return -1
	}
	x := t.root
	for ; x != nil && x.children[0] != nil; height++ {
		x = x.children[0]
	}
	return
}

// Empty 判空
func (t *Btree) Empty() bool {
	return t.size <= 0
}

// Order 当前阶数
func (t *Btree) Order() int {
	return t.order
}

// Clear 清空B-树
func (t *Btree) Clear() int {
	size := t.size
	t.hot = nil
	t.root = nil
	t.size = 0
	return size
}

// Search 元素查找
func (t *Btree) Search(item interface{}) interface{} {
	node, i := t.searchAt(t.root, item)
	if node == nil {
		return nil
	}
	return node.items[i]
}

// Insert 元素新增：
// 算法：
// 1、找到元素 x 对于节点，存在则返回。
// 2、令最后访问节点为 x， 在 x 中查找元素适合的插入位置 i， 并将其插入。
// 4、x上溢修复。
func (t *Btree) Insert(item interface{}) {
	if t.Empty() {
		t.size++
		t.root = emptyBtreeNode(nil)
		t.root.insertItemAt(0, item)
		t.root.insertChildAt(1, nil)
		return
	}
	node, i := t.searchAt(t.root, item)
	if node != nil {
		node.items[i] = item
		return
	}
	t.size++
	node = t.hot
	i, _ = node.find(t.keyComparator, item)
	node.insertItemAt(i, item)
	node.insertChildAt(i+1, nil)
	t.solveOverflow(node)
}

// Remove 元素删除：
// 算法：
// 1、找到元素 x 对于节点。
// 2、若节点 x 不为叶节点，则交换 x 与其后继节点。
// 3、在x中删除对于元素。
// 4、在x下溢修复。
func (t *Btree) Remove(item interface{}) {
	node, i := t.searchAt(t.root, item)
	if node == nil {
		return
	}
	t.size--
	if node.children[0] != nil {
		suc := node.children[i+1]
		for suc.children[0] != nil {
			suc = suc.children[0]
		}
		node.items[i], suc.items[0] = suc.items[0], node.items[i]
		node, i = suc, 0
	}
	node.removeItemAt(i)
	node.removeChildAt(i + 1)
	t.solveUnderflow(node)
}

// 层序遍历打印当前树中节点信息
func (t *Btree) levelInfo() string {
	var info []string
	if t.Empty() {
		return "[]"
	}
	que := queue.New()
	que.Push(t.root)
	for !que.Empty() {
		var tmp string
		size := que.Size()
		for ; size > 0; size-- {
			node := que.Pop().(*btreeNode)
			tmp += fmt.Sprintf("%v ", node.items)
			if node.children[0] != nil {
				for _, child := range node.children {
					que.Push(child)
				}
			}
		}
		tmp = strings.TrimRight(tmp, " ")
		info = append(info, tmp)
	}
	return "[" + strings.Join(info, "\n") + "]"
}
