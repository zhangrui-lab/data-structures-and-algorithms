// 向量
package vector

import (
	"math/rand"
)

type Rank int

const defaultCapacity = 3

type vector struct {
	elem []interface{}
	size Rank
	cap  Rank
}

func New() *vector {
	return &vector{
		elem: make([]interface{}, defaultCapacity, defaultCapacity),
		size: 0,
		cap:  defaultCapacity,
	}
}

// 返回向量的大小
func (vec *vector) Size() int {
	return int(vec.size)
}

// 判定向量是否为空
func (vec *vector) Empty() bool {
	return vec.size <= 0
}

// 返回向量的容量
func (vec *vector) Capacity() int {
	return int(vec.cap)
}

// 返回向量中第一个元素。
func (vec *vector) Front() interface{} {
	if vec.Empty() {
		return nil
	}
	return vec.elem[0]
}

// 返回向量中最后一个元素。
func (vec *vector) Back() interface{} {
	if vec.Empty() {
		return nil
	}
	return vec.elem[vec.size-1]
}

// 返回向量中位置 n 处元素。
func (vec *vector) At(n int) interface{} {
	if !vec.validRank(n) {
		return nil
	}
	return vec.elem[n]
}

// 为向量分配新内容，替换其当前内容，并相应地修改其大小。
func (vec *vector) Assign(n int, val interface{}) bool {
	if !vec.validRank(n) {
		return false
	}
	vec.elem[n] = val
	return true
}

// 在[0,size)的指定位置进行插入：插入位置大于容器当前元素数时，插入到末尾
func (vec *vector) insert(n int, val interface{}) bool {
	if !vec.validRank(n) {
		return false
	}
	r := Rank(n)
	vec.expand()
	for i := vec.size; i > r; i-- {
		vec.elem[i] = vec.elem[i-1]
	}
	vec.size++
	vec.elem[r] = val
	return true
}

// 从向量中移除单个元素(位置)
func (vec *vector) Remove(n int) interface{} {
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
	for ; Rank(hi) < vec.size; lo++ {
		vec.elem[lo] = vec.elem[hi]
		hi++
	}
	vec.size = Rank(lo)
	vec.shrink()
	return hi - lo
}

// 尾部进行插入
func (vec *vector) Push(val interface{}) {
	vec.expand()
	vec.elem[vec.size] = val
	vec.size++
}

// 尾部进行删除
func (vec *vector) Pop() interface{} {
	return vec.Remove(int(vec.size) - 1)
}

// 无序向量的查找：失败时返回-1
func (vec *vector) Find(val interface{}) int {
	return vec.FindRange(0, int(vec.size), val)
}

// 无序向量区间[lo, hi)的查找：失败时返回-1
func (vec *vector) FindRange(lo, hi int, val interface{}) int {
	for ; lo < hi && vec.elem[hi-1] == val; hi-- {
	}
	return hi
}

// 整体置乱
func (vec *vector) Scrambling() {
	vec.ScramblingRange(0, int(vec.size))
}

// 对区间[lo, hi)置乱
func (vec *vector) ScramblingRange(lo, hi int) {
	for ; lo < hi; hi-- {
		pos := rand.Intn(int(vec.size))
		vec.elem[pos], vec.elem[hi-1] = vec.elem[hi-1], vec.elem[pos]
	}
}

// 元素遍历
func (vec *vector) Traverse(visit func(*interface{})) {
	for i := Rank(0); i < vec.size; i++ {
		visit(&vec.elem[i])
	}
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
	newElem := make([]interface{}, vec.cap, vec.cap)
	copy(newElem, vec.elem)
	vec.elem = newElem
}

// 装填因子过小时压缩
func (vec *vector) shrink() {
	if vec.size<<2 > vec.cap || vec.cap < defaultCapacity<<1 {
		return
	}
	vec.cap >>= 1
	newElem := make([]interface{}, vec.cap, vec.cap)
	copy(newElem, vec.elem)
	vec.elem = newElem
}

// 验证秩是否合法
func (vec *vector) validRank(n int) bool {
	return 0 <= n && Rank(n) < vec.size
}
