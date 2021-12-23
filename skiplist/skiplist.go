package skiplist

import (
	"data-structures-and-algorithms/contract"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	maxLevel int = 3
)

// quadListNode 四链表节点
// 	     ^
// 	     |
// 	<-- node -->
// 	     |
// 	     v
type quadListNode struct {
	key          interface{}
	value        interface{}
	pred, next   *quadListNode //前驱、后继
	above, below *quadListNode //上邻、下邻
}

// 从链接关系中移除当前节点
func (q *quadListNode) remove() {
	q.pred.next = q.next
	if q.next != nil {
		q.next.pred = q.pred
	}
	if q.above != nil {
		q.above.below = q.below
	}
	if q.below != nil {
		q.below.above = q.above
	}
	q.pred = nil
	q.next = nil
	q.above = nil
	q.below = nil
}

// SkipList 跳转表结构
type SkipList struct {
	size       int
	header     []*quadListNode
	maxLevel   int
	comparator contract.Comparator
	hot        *quadListNode
}

// NewSkipList 新建空跳表
func NewSkipList(cmps ...contract.Comparator) *SkipList {
	cmp := contract.DefaultComparator
	if len(cmps) > 0 {
		cmp = cmps[0]
	}
	rand.Seed(time.Now().UnixNano())
	var header []*quadListNode
	for i := 0; i < maxLevel; i++ {
		header = append(header, &quadListNode{})
		if i > 0 {
			header[i-1].above = header[i]
			header[i].below = header[i-1]
		}
	}
	return &SkipList{
		header:     header,
		maxLevel:   maxLevel,
		comparator: cmp,
	}
}
func (s *SkipList) Size() int {
	return s.size
}
func (s *SkipList) Empty() bool {
	return s.size <= 0
}

func (s *SkipList) Search(key interface{}) interface{} {
	q := s.search(key)
	if q == nil {
		return nil
	}
	return q.value
}

func (s *SkipList) Insert(key, value interface{}) {
	if s.Empty() {
		s.size++
		s.header[0].next = &quadListNode{key: key, value: value, pred: s.header[0]}
		return
	}
	q := s.search(key)
	if q != nil { // 存在并替换
		for q != nil {
			q.value = value
			q = q.below
		}
		return
	}
	s.size++
	hot := s.hot
	var pred, next *quadListNode
	if s.comparator(hot.key, key) <= 0 {
		pred, next = hot, hot.next
	} else {
		pred, next = hot.pred, hot
	}
	q = &quadListNode{key: key, value: value, pred: pred, next: next}
	q.pred.next = q
	if q.next != nil {
		q.next.pred = q
	}
	s.towerGrowth(q)
}

// Remove 删除key
func (s *SkipList) Remove(key interface{}) {
	if s.Empty() {
		return
	}
	q := s.search(key)
	if q == nil {
		return
	}
	s.towerRemove(q)
}

// 移除以 q 为底的塔
func (s *SkipList) towerRemove(q *quadListNode) {
	for q != nil {
		a := q.above
		q.remove()
		q = a
	}
}

// 增长以 q 为底的塔
func (s *SkipList) towerGrowth(q *quadListNode) {
	for num := rand.Intn(40); num >= 3; num = rand.Intn(10) {
		pred := q.pred.above
		if pred == nil {
			break
		}
		next := pred.next
		node := &quadListNode{key: q.key, value: q.value, pred: pred, next: next, below: q}
		q.above = node
		pred.next = node
		if next != nil {
			next.pred = node
		}
		q = node
	}
}

// 在跳转表中进行搜索，找到对应元素的最高层次节点：若找到元素，则返回该元素塔底指针，否则返回该最后访问的塔底元素。
func (s *SkipList) search(key interface{}) *quadListNode {
	var q = s.header[maxLevel-1]
	for q.next == nil {
		q = q.below
	}
	q = q.next
	for q != nil && s.comparator(q.key, key) != 0 {
		for s.hot = q; q != nil && s.comparator(q.key, key) < 0; q, s.hot = q.next, q {
		}
		if q == nil || (s.comparator(q.key, key) > 0 && s.hot != q) {
			q = s.hot.below
			continue
		}
		if s.comparator(q.key, key) > 0 {
			q = q.pred.below
			if q != nil {
				q = q.next
			}
		}
	}
	for q != nil && q.below != nil {
		q = q.below
	}
	return q
}

func (s *SkipList) levelInfo() string {
	var info []string
	if s.Empty() {
		return ""
	}
	var q = s.header[maxLevel-1]
	for q.next == nil {
		q = q.below
	}
	for ; q != nil; q = q.below {
		tmp := "{"
		p := q.next
		for p != nil {
			tmp += fmt.Sprintf("%v,", p.key)
			p = p.next
		}
		tmp = strings.TrimRight(tmp, ",")
		tmp += "}"
		info = append(info, tmp)
	}
	return strings.Join(info, "\n")
}
