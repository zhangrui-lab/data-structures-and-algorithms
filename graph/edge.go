package graph

type EdgeType uint8

const (
	_            EdgeType = iota
	UNDETERMINED          // 初始化状态，待定
	TREE                  // 树边：图访问中将某一节点引入支撑树的第一条边
	CROSS                 // 跨边
	FORWARD               // 前向边：dfs中由支撑树的未访问完成的父节点指向已访问完成的子节点的边
	BACKWARD              // 后向边：dfs中由支撑树的已访问完成的子节点指向未访问完成的父节点的边（环路存在）
)

// 邻接矩阵树边对象
type edge struct {
	weight interface{} // 权重信息.可在边中保存数据信息，并一句需要的属性进行权重比较
	genre  EdgeType
}

// 以权重weight新建树边
func newEdge(weight interface{}) *edge {
	return &edge{weight: weight, genre: UNDETERMINED}
}

// 边 e 辅助信息复位
func (e *edge) reset() {
	e.genre = UNDETERMINED
}

// 邻接表树边对象
type adjEdge struct {
	node *vertex // 指向的节点
	edge         // 边信息
}
