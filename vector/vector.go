// 向量
package vector

import (
	"math/rand"
)

const defaultCapacity = 3

type vector struct {
	elem []Sortable
	size int
	cap  int
}

// 直接实例化
func New() *vector {
	return &vector{
		elem: make([]Sortable, defaultCapacity, defaultCapacity),
		size: 0,
		cap:  defaultCapacity,
	}
}

// 复制另一向量
func Copy(o *vector) *vector {
	vec := New()
	vec.elem = make([]Sortable, o.size, o.cap)
	copy(vec.elem, o.elem)
	vec.size = o.size
	vec.cap = o.cap
	return vec
}

// 从切片中复制
func CopySlice(o []Sortable) *vector {
	vec := New()
	vec.elem = make([]Sortable, len(o), cap(o))
	copy(vec.elem, o)
	vec.size = len(o)
	vec.cap = cap(o)
	return vec
}

// 返回向量的大小
func (vec *vector) Size() int {
	return vec.size
}

// 返回vec[i]与vec[j]的大小
func (vec *vector) Less(i, j int) bool {
	return vec.elem[i].Less(vec.elem[j])
}

// 交换i，j 元素位置
func (vec *vector) Swap(i, j int) {
	vec.elem[i], vec.elem[j] = vec.elem[j], vec.elem[i]
}

// 判定向量是否为空
func (vec *vector) Empty() bool {
	return vec.size <= 0
}

// 返回向量的容量
func (vec *vector) Capacity() int {
	return vec.cap
}

// 返回向量中第一个元素。
func (vec *vector) Front() Sortable {
	if vec.Empty() {
		return nil
	}
	return vec.elem[0]
}

// 返回向量中最后一个元素。
func (vec *vector) Back() Sortable {
	if vec.Empty() {
		return nil
	}
	return vec.elem[vec.size-1]
}

// 返回向量中位置 n 处元素。
func (vec *vector) At(n int) Sortable {
	if !vec.validRank(n) {
		return nil
	}
	return vec.elem[n]
}

// 为向量分配新内容，替换其当前内容，并相应地修改其大小。
func (vec *vector) Assign(n int, v Sortable) bool {
	if !vec.validRank(n) {
		return false
	}
	vec.elem[n] = v
	return true
}

// 在[0,size)的指定位置进行插入：插入位置大于容器当前元素数时，插入到末尾
func (vec *vector) insert(n int, v Sortable) bool {
	if !vec.validRank(n) {
		return false
	}
	r := n
	vec.expand()
	for i := vec.size; i > r; i-- {
		vec.elem[i] = vec.elem[i-1]
	}
	vec.size++
	vec.elem[r] = v
	return true
}

// 从向量中移除单个元素(位置)
func (vec *vector) Remove(n int) Sortable {
	if !vec.validRank(n) {
		return nil
	}
	e := vec.elem[n]
	vec.RemoveRange(n, n+1)
	return e
}

// 从向量中移除一系列元素([lo,hi)) 并返回移除元素个数
func (vec *vector) RemoveRange(lo, hi int) int {
	if lo >= hi || !vec.validRank(lo) {
		return 0
	}
	for ; hi < vec.size; lo++ {
		vec.elem[lo] = vec.elem[hi]
		hi++
	}
	vec.size = lo
	vec.shrink()
	return hi - lo
}

// 清空向量
func (vec *vector) Clear() int {
	sz := vec.size
	vec.elem = vec.elem[:0]
	return sz
}

// 尾部进行插入
func (vec *vector) Push(v Sortable) {
	vec.expand()
	vec.elem[vec.size] = v
	vec.size++
}

// 尾部进行删除
func (vec *vector) Pop() Sortable {
	return vec.Remove(vec.size - 1)
}

// 无序向量的查找：失败时返回-1
func (vec *vector) Find(v Sortable) int {
	return vec.FindRange(v, 0, vec.size)
}

// 无序向量区间[lo, hi)的查找：失败时返回-1
func (vec *vector) FindRange(v Sortable, lo, hi int) int {
	for ; lo < hi && vec.elem[hi-1] == v; hi-- {
	}
	return hi
}

// 整体置乱
func (vec *vector) Scrambling() {
	vec.ScramblingRange(0, vec.size)
}

// 对区间[lo, hi)置乱
func (vec *vector) ScramblingRange(lo, hi int) {
	for ; lo < hi; hi-- {
		pos := rand.Intn(vec.size)
		vec.elem[pos], vec.elem[hi-1] = vec.elem[hi-1], vec.elem[pos]
	}
}

// 元素遍历
func (vec *vector) Traverse(visit func(*Sortable)) {
	for i := 0; i < vec.size; i++ {
		visit(&vec.elem[i])
	}
}

// 返回向量的逆序对数
func (vec *vector) disordered() int {
	var num = 0
	for i := 1; i < vec.size; i++ {
		if vec.elem[i].Less(vec.elem[i-1]) {
			num++
		}
	}
	return num
}

// 删除无序向量中重复元素（靠后者）
func (vec *vector) deduplicate() int {
	oldSize := vec.size
	set := make(map[Sortable]struct{})
	for i := 0; i < vec.size; i++ {
		if _, ok := set[vec.elem[i]]; ok {
			vec.Remove(i)
		} else {
			set[vec.elem[i]] = struct{}{}
		}
	}
	vec.shrink()
	return vec.size - oldSize
}

//有序去重
func (vec *vector) uniquify() int {
	var i, j = 0, 1
	for ; j < vec.size; j++ {
		if vec.elem[i] != vec.elem[j] {
			i++
			vec.elem[i] = vec.elem[j]
		}
	}
	vec.size = i + 1
	vec.shrink()
	return j - i
}

// 有序向量整体查找
func (vec *vector) Search(v Sortable) int {
	return vec.SearchRange(v, 0, vec.size)
}

// 有序向量区间[lo, hi)查找
func (vec *vector) SearchRange(v Sortable, lo, hi int) int {
	return vec.binarySearchV2(v, lo, hi)
}

// 朴素二分查找
func (vec *vector) binarySearchV1(v Sortable, lo, hi int) int {
	for lo < hi {
		mid := (lo + hi) >> 1
		if v.Less(vec.elem[mid]) {
			hi = mid
		} else if vec.elem[mid].Less(v) {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

// 优化二分查找
func (vec *vector) binarySearchV2(v Sortable, lo, hi int) int {
	for lo < hi {
		mid := (lo + hi) >> 1
		if v.Less(vec.elem[mid]) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo - 1
}

// 空间不足时扩容
func (vec *vector) expand() {
	if vec.size < vec.cap {
		return
	}
	if vec.cap < defaultCapacity {
		vec.cap = defaultCapacity
	}
	vec.cap <<= 1
	newElem := make([]Sortable, vec.cap, vec.cap)
	copy(newElem, vec.elem)
	vec.elem = newElem
}

// 装填因子过小时压缩：空间利用率保持在 50% 之上
func (vec *vector) shrink() {
	if vec.size<<2 > vec.cap || vec.cap < defaultCapacity<<1 {
		return
	}
	vec.cap >>= 1
	newElem := make([]Sortable, vec.cap, vec.cap)
	copy(newElem, vec.elem)
	vec.elem = newElem
}

// 验证秩是否合法
func (vec *vector) validRank(n int) bool {
	return 0 <= n && n < vec.size
}
