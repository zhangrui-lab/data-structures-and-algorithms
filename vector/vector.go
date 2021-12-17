// Package vector 向量
package vector

import (
	"data-structures-and-algorithms/types"
	"fmt"
	"math/rand"
	"strings"
)

const lowCapacity = 3

type vector struct {
	elem []types.Sortable
	size int
	cap  int
}

// New 直接实例化
func New(cap int) *vector {
	return &vector{
		elem: make([]types.Sortable, 0, cap),
		size: 0,
		cap:  cap,
	}
}

// Copy 复制另一向量
func Copy(o *vector) *vector {
	vec := New(o.cap << 1)
	vec.elem = make([]types.Sortable, o.size, o.cap<<1)
	copy(vec.elem, o.elem)
	vec.size = o.size
	return vec
}

// CopySlice 复制切片以创建向量
func CopySlice(o ...types.Sortable) *vector {
	vec := New(cap(o) << 1)
	vec.elem = make([]types.Sortable, len(o), cap(o)<<1)
	copy(vec.elem, o)
	vec.size = len(o)
	return vec
}

// Less 返回vec[i]与vec[j]的大小
func (vec *vector) Less(i, j int) bool {
	return vec.elem[i].Less(vec.elem[j])
}

// Swap 交换i，j 元素位置
func (vec *vector) Swap(i, j int) {
	vec.elem[i], vec.elem[j] = vec.elem[j], vec.elem[i]
}

// Size 返回向量的大小
func (vec *vector) Size() int {
	return vec.size
}

// Empty 判定向量是否为空
func (vec *vector) Empty() bool {
	return vec.size <= 0
}

// Capacity 返回向量的容量
func (vec *vector) Capacity() int {
	return vec.cap
}

// Front 返回向量中第一个元素。
func (vec *vector) Front() (types.Sortable, error) {
	if vec.Empty() {
		return nil, fmt.Errorf("vector is empty")
	}
	return vec.elem[0], nil
}

// Back 返回向量中最后一个元素。
func (vec *vector) Back() (types.Sortable, error) {
	if vec.Empty() {
		return nil, fmt.Errorf("vector is empty")
	}
	return vec.elem[vec.size-1], nil
}

// At 返回向量中位置 n 处元素。
func (vec *vector) At(r int) (types.Sortable, error) {
	if !vec.validRank(r) {
		return nil, fmt.Errorf("access out of bounds. len = %v, idx = %v", vec.size, r)
	}
	return vec.elem[r], nil
}

// Assign 为向量分配新内容，替换其当前内容。 当 n 非法时返回false。
func (vec *vector) Assign(r int, v types.Sortable) bool {
	if !vec.validRank(r) {
		return false
	}
	vec.elem[r] = v
	return true
}

// Disordered 返回向量的逆序对数
func (vec *vector) Disordered() int {
	var num = 0
	for i := 1; i < vec.size; i++ {
		if vec.elem[i].Less(vec.elem[i-1]) {
			num++
		}
	}
	return num
}

// String 字符串形式
func (vec *vector) String() string {
	if vec.Empty() {
		return "{}"
	}
	items := make([]string, 0, vec.size)
	for _, item := range vec.elem {
		items = append(items, fmt.Sprintf("%v", item))
	}
	return "{" + strings.Join(items, ", ") + "}"
}

// insert 在 [0,size] 的指定位置处进行插入。 当 n 非法时返回false。
func (vec *vector) insert(r int, v types.Sortable) bool {
	if r != vec.size && !vec.validRank(r) {
		return false
	}
	vec.expand()
	vec.elem = vec.elem[:vec.size+1]
	for i := vec.size; i > r; i-- {
		vec.elem[i] = vec.elem[i-1]
	}
	vec.size++
	vec.elem[r] = v
	return true
}

// Remove 移除向量中秩为 r 的元素
func (vec *vector) Remove(r int) (types.Sortable, error) {
	if !vec.validRank(r) {
		return nil, fmt.Errorf("out of bounds. len = %v, idx = %v", vec.size, r)
	}
	e := vec.elem[r]
	vec.RemoveRange(r, r+1)
	return e, nil
}

// RemoveRange 移除秩在区间 [lo,hi) 中的元素
func (vec *vector) RemoveRange(lo, hi int) {
	if lo >= hi || !vec.validRank(lo) {
		return
	}
	vec.elem = append(vec.elem[:lo], vec.elem[hi:]...)
	vec.size = len(vec.elem)
	vec.shrink()
}

// Clear 清空向量，不收缩所占空间
func (vec *vector) Clear() {
	vec.size = 0
	vec.elem = vec.elem[:0]
}

// Push 尾部进行插入
func (vec *vector) Push(v types.Sortable) {
	vec.insert(vec.size, v)
}

// Pop 尾部进行删除
func (vec *vector) Pop() (types.Sortable, error) {
	return vec.Remove(vec.size - 1)
}

// Scrambling 向量整体置乱
func (vec *vector) Scrambling() {
	vec.ScramblingRange(0, vec.size)
}

// ScramblingRange 向量区间[lo, hi)置乱
func (vec *vector) ScramblingRange(lo, hi int) {
	for ; lo < hi; hi-- {
		vec.Swap(rand.Intn(vec.size), hi-1)
	}
}

// Traverse 遍历向量元素
func (vec *vector) Traverse(visit func(*types.Sortable)) {
	for i := 0; i < vec.size; i++ {
		visit(&vec.elem[i])
	}
}

// Find 无序向量查找：多个元素时返回秩最大者，失败时返回-1
func (vec *vector) Find(v types.Sortable) int {
	return vec.FindRange(v, 0, vec.size)
}

// FindRange 无序向量区间 [lo, hi) 查找：失败时返回-1
func (vec *vector) FindRange(v types.Sortable, lo, hi int) int {
	for ; lo < hi && vec.elem[hi-1] != v; hi-- {
	}
	if lo == hi {
		return -1
	}
	return hi - 1
}

// Deduplicate 无序向量去重
func (vec *vector) Deduplicate() int {
	oldSize := vec.size
	set := make(map[types.Sortable]struct{})
	for i := 0; i < vec.size; i++ {
		if _, ok := set[vec.elem[i]]; ok {
			_, _ = vec.Remove(i)
		} else {
			set[vec.elem[i]] = struct{}{}
		}
	}
	return vec.size - oldSize
}

// Search 有序向量整体查找, 返回不大于v的元素的最大秩
func (vec *vector) Search(v types.Sortable) int {
	return vec.SearchRange(v, 0, vec.size)
}

// SearchRange 有序向量区间 [lo, hi) 查找, 返回不大于v的元素的最大秩
func (vec *vector) SearchRange(v types.Sortable, lo, hi int) int {
	return vec.binarySearchV2(v, lo, hi)
}

// Uniquify 有序向量去重
func (vec *vector) Uniquify() int {
	var i, j = 0, 1
	for ; j < vec.size; j++ {
		if vec.elem[i] != vec.elem[j] {
			i++
			vec.elem[i] = vec.elem[j]
		}
	}
	vec.size = i + 1
	vec.elem = vec.elem[:vec.size]
	vec.shrink()
	return j - vec.size
}

// 朴素二分查找
func (vec *vector) binarySearchV1(v types.Sortable, lo, hi int) int {
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
func (vec *vector) binarySearchV2(v types.Sortable, lo, hi int) int {
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

// 扩容：空间不足时对容量执行翻倍
func (vec *vector) expand() {
	if vec.size < vec.cap {
		return
	}
	if vec.cap < lowCapacity {
		vec.cap = lowCapacity
	}
	vec.cap <<= 1
	newElem := make([]types.Sortable, vec.size, vec.cap)
	copy(newElem, vec.elem)
	vec.elem = newElem
}

// 缩容：维持空间利用率在 50% 之上
func (vec *vector) shrink() {
	if vec.size<<2 > vec.cap || vec.cap < lowCapacity<<1 {
		return
	}
	vec.cap >>= 1
	newElem := make([]types.Sortable, vec.size, vec.cap)
	copy(newElem, vec.elem)
	vec.elem = newElem
}

// 验证秩是否合法
func (vec *vector) validRank(r int) bool {
	return 0 <= r && r < vec.size
}
