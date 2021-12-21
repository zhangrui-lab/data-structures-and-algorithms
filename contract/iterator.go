package contract

// Iterator 前向迭代器
type Iterator interface {
	Valid() bool
	Value() interface{}
	Set(interface{})
	Next() Iterator // 直接后继
	Equal(o Iterator) bool
}

// BidIterator 双向迭代器
type BidIterator interface {
	Prev() BidIterator // 直接前驱
}

// KvIterator 键值对前向迭代器
type KvIterator interface {
	Key() interface{}
}

// KvBidIterator 键值对双向迭代器
type KvBidIterator interface {
	Prev() KvBidIterator
}

// RandomAccessIterator 随机访问迭代器
type RandomAccessIterator interface {
	BidIterator
	At(position int) RandomAccessIterator // 寻秩
	forward(n int) RandomAccessIterator   // 向前
	backward(n int) RandomAccessIterator  // 向后
}

// Swap 交换迭代器值元素
func Swap(iter1, iter2 Iterator) {
	v1 := iter1.Valid()
	iter1.Set(iter2.Value())
	iter2.Set(v1)
}
