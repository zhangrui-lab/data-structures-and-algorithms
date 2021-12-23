package priority_queue

type PQ interface {
	// Insert 按照比较器确定的优先级次序插入词条
	Insert(item interface{})
	// GetMax 取出优先级最高的词条
	GetMax() interface{}
	// DelMax 删除优先级最高的词条
	DelMax()
}
