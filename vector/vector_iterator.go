package vector

import "data-structures-and-algorithms/contract"

var _ contract.RandomAccessIterator = (*VectorIterator)(nil)

// VectorIterator 向量迭代器
type VectorIterator struct {
	vector *Vector
	rank   int
}

// Valid 是否为合法迭代器
func (iter *VectorIterator) Valid() bool {
	return iter != nil && iter.vector.validRank(iter.rank)
}

// Value 取值
func (iter *VectorIterator) Value() interface{} {
	return iter.vector.At(iter.rank)
}

// Set 设置值
func (iter *VectorIterator) Set(value interface{}) {
	iter.vector.Assign(iter.rank, value)
}

// Next 直接后继
func (iter *VectorIterator) Next() contract.Iterator {
	iter.rank++
	if !iter.vector.validRank(iter.rank) {
		iter.rank = iter.vector.Size()
	}
	return iter
}

// Prev 直接前驱
func (iter *VectorIterator) Prev() contract.BidIterator {
	iter.rank--
	return iter
}

// Forward 从当前位置向前 n，不合法时返回nil
func (iter *VectorIterator) Forward(n int) contract.RandomAccessIterator {
	iter.rank += n
	if !iter.vector.validRank(iter.rank) {
		iter.rank = iter.vector.Size()
	}
	return iter
}

// Backward 从当前位置向后 n，不合法时返回nil
func (iter *VectorIterator) Backward(n int) contract.RandomAccessIterator {
	iter.rank -= n
	return iter
}

// At 秩 n 处的迭代器：n 不合法时返回nil
func (iter *VectorIterator) At(n int) contract.RandomAccessIterator {
	iter.rank = n
	if n >= iter.vector.Size() {
		iter.rank = iter.vector.Size()
	}
	return iter
}

// Equal 迭代器判等
func (iter *VectorIterator) Equal(other contract.Iterator) bool {
	otherIter, ok := other.(*VectorIterator)
	if !ok {
		return false
	}
	if otherIter.vector == iter.vector && otherIter.rank == iter.rank {
		return true
	}
	return false
}
