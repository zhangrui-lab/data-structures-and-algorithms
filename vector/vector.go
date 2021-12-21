// Package vector 向量
package vector

import (
	"data-structures-and-algorithms/contract"
	"fmt"
	"math/rand"
	"strings"
)

const lowCapacity = 3

type Vector struct {
	elem []interface{}
}

// New 直接实例化
func New(cap int) *Vector {
	return &Vector{
		elem: make([]interface{}, 0, cap),
	}
}

// Copy 复制另一向量
func Copy(o *Vector) *Vector {
	vec := New(o.Capacity())
	vec.elem = vec.elem[:o.Size()]
	copy(vec.elem, o.elem)
	return vec
}

// FromSlice 复制切片以创建向量
func FromSlice(o ...interface{}) *Vector {
	vec := New(cap(o))
	vec.elem = vec.elem[:len(o)]
	copy(vec.elem, o)
	return vec
}

// Size 返回向量的大小
func (vec *Vector) Size() int {
	return len(vec.elem)
}

// Empty 判定向量是否为空
func (vec *Vector) Empty() bool {
	return vec.Size() <= 0
}

// Capacity 返回向量的容量
func (vec *Vector) Capacity() int {
	return cap(vec.elem)
}

// Front 返回向量中第一个元素。
func (vec *Vector) Front() interface{} {
	if vec.Empty() {
		return nil
	}
	return vec.elem[0]
}

// Back 返回向量中最后一个元素。
func (vec *Vector) Back() interface{} {
	if vec.Empty() {
		return nil
	}
	return vec.elem[vec.Size()-1]
}

// First 首元素迭代器
func (vec *Vector) First() *VectorIterator {
	return vec.Begin()
}

// Last 末元素迭代器
func (vec *Vector) Last() *VectorIterator {
	return &VectorIterator{vec, vec.Size() - 1}
}

// Begin 首元素迭代器
func (vec *Vector) Begin() *VectorIterator {
	return &VectorIterator{vec, 0}
}

// End 尾后元素迭代器·
func (vec *Vector) End() *VectorIterator {
	return &VectorIterator{vector: vec, rank: vec.Size()}
}

// At 返回向量中位置 i 处元素。
func (vec *Vector) At(i int) interface{} {
	if !vec.validRank(i) {
		return nil
	}
	return vec.elem[i]
}

// Assign 为向量指定合法位置 r 分配新内容，替换其当前内容。
func (vec *Vector) Assign(r int, v interface{}) {
	if !vec.validRank(r) {
		return
	}
	vec.elem[r] = v
}

// String 字符串形式
func (vec *Vector) String() string {
	if vec.Empty() {
		return "{}"
	}
	items := make([]string, 0, vec.Size())
	for _, item := range vec.elem {
		items = append(items, fmt.Sprintf("%v", item))
	}
	return "{" + strings.Join(items, ", ") + "}"
}

// insert 在 [0,size] 的指定位置处进行插入。 当 n 非法时返回false。
func (vec *Vector) insert(r int, v interface{}) bool {
	size := vec.Size()
	if r != size && !vec.validRank(r) {
		return false
	}
	vec.expand()
	vec.elem = vec.elem[:size+1]
	for i := size; i > r; i-- {
		vec.elem[i] = vec.elem[i-1]
	}
	vec.elem[r] = v
	return true
}

// Remove 移除向量中秩为 r 的元素
func (vec *Vector) Remove(r int) interface{} {
	if !vec.validRank(r) {
		return nil
	}
	e := vec.elem[r]
	vec.RemoveRange(r, r+1)
	return e
}

// RemoveRange 移除秩在区间 [lo,hi) 中的元素
func (vec *Vector) RemoveRange(lo, hi int) {
	if lo >= hi || !vec.validRank(lo) {
		return
	}
	vec.elem = append(vec.elem[:lo], vec.elem[hi:]...)
	vec.shrink()
}

// Clear 清空向量，不收缩所占空间
func (vec *Vector) Clear() {
	vec.elem = vec.elem[:0]
}

// PushBack 尾部进行插入
func (vec *Vector) PushBack(value interface{}) {
	vec.insert(vec.Size(), value)
}

// PopBack 尾部进行删除
func (vec *Vector) PopBack() interface{} {
	return vec.Remove(vec.Size() - 1)
}

// PushFront 首部进行插入：O(n) 时间复杂度
func (vec *Vector) PushFront(value interface{}) {
	vec.insert(0, value)
}

// PopFront 首部进行删除：O(n) 时间复杂度
func (vec *Vector) PopFront() interface{} {
	return vec.Remove(0)
}

// Scrambling 向量整体置乱
func (vec *Vector) Scrambling() {
	vec.ScramblingRange(0, vec.Size())
}

// ScramblingRange 向量区间[lo, hi)置乱
func (vec *Vector) ScramblingRange(lo, hi int) {
	for ; lo < hi; hi-- {
		i := rand.Intn(hi)
		vec.elem[hi-1], vec.elem[i] = vec.elem[i], vec.elem[hi-1]
	}
}

// Traverse 遍历向量元素
func (vec *Vector) Traverse(visitor contract.Visitor) {
	for i := 0; i < vec.Size(); i++ {
		visitor(vec.elem[i])
	}
}

// Deduplicate 无序向量去重
func (vec *Vector) Deduplicate() int {
	oldSize := vec.Size()
	set := make(map[interface{}]struct{})
	for i := 0; i < vec.Size(); i++ {
		if _, ok := set[vec.elem[i]]; ok {
			vec.Remove(i)
		} else {
			set[vec.elem[i]] = struct{}{}
		}
	}
	return oldSize - vec.Size()
}

// Uniquify 有序向量去重
func (vec *Vector) Uniquify() int {
	var i, j = 0, 1
	for ; j < vec.Size(); j++ {
		if vec.elem[i] != vec.elem[j] {
			i++
			vec.elem[i] = vec.elem[j]
		}
	}
	vec.elem = vec.elem[:i+1]
	vec.shrink()
	return j - vec.Size()
}

// 扩容：空间不足时对容量执行翻倍
func (vec *Vector) expand() {
	if vec.Size() < vec.Capacity() {
		return
	}
	cap := vec.Capacity()
	if cap < lowCapacity {
		cap = lowCapacity
	} else {
		cap <<= 1
	}
	newElem := make([]interface{}, vec.Size(), cap)
	copy(newElem, vec.elem)
	vec.elem = newElem
}

// 缩容：维持空间利用率在 50% 之上
func (vec *Vector) shrink() {
	if vec.Size()<<2 > vec.Capacity() || vec.Capacity() < lowCapacity<<1 {
		return
	}
	cap := vec.Capacity() >> 1
	newElem := make([]interface{}, vec.Size(), cap)
	copy(newElem, vec.elem)
	vec.elem = newElem
}

// 验证秩是否合法
func (vec *Vector) validRank(r int) bool {
	return 0 <= r && r < vec.Size()
}
