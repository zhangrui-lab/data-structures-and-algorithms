package contract

// SeqContainer 线性容器通用接口
type SeqContainer interface {
	Empty() bool
	Size() int
	String() string
	Front() interface{}
	Back() interface{}
	Begin() Iterator
	End() Iterator
	At(i int) interface{}

	Clear()
	PushFront(value interface{})
	PopFront() interface{}
	PushBack(value interface{})
	PopBack() interface{}
	Swap(o SeqContainer)
}
