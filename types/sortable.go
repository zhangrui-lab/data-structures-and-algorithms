package types

// 比较器
type Sortable interface {
	// 对两个课比较对象进行比较， 当 x < y 时， x.less(y) 返回 true
	// 当 x.less(y) == false && y.less(x) == false 时，x == y
	Less(o Sortable) bool
}

// 可排序对象接口
type Interface interface {
	Size() int
	Less(i, j int) bool
	Swap(i, j int)
}
