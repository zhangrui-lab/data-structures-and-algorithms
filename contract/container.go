package contract

// SeqContainer 线性容器通用接口
type SeqContainer interface {
	Size() int
	Empty() bool
	Front() interface{}
	Back() interface{}
	At(i int) interface{}
	String() string

	Clear()
	PushFront(value interface{})
	PopFront() interface{}
	PushBack(value interface{})
	PopBack() interface{}
}
