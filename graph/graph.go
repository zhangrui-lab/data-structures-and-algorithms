package graph

// Graph 图结构接口约束
type Graph interface {
	// NodeNumbers 节点总数
	NodeNumbers() int
	// Exists 节点 v 是否存在
	Exists(v int) bool
	// Insert 插入顶点，返回编号
	Insert(data interface{}) int
	// Remove 删除顶点v及其关联边，返回该顶点信息
	Remove(v int) interface{}
	// Vertex 顶点v的数据
	Vertex(v int) interface{}
	// InDegree 顶点v的入度
	InDegree(v int) int
	// OutDegree 顶点v的出度
	OutDegree(v int) int
	// FirstNbr 顶点v的首个邻接顶点
	FirstNbr(v int) int
	// NextNbr 顶点v的, 相对于顶点i的下一邻接顶点
	NextNbr(v int, i int) int
	// Status 顶点v的状态
	Status(v int) NodeState
	// DTime 顶点v的时间标签dTime
	DTime(v int) int
	// FTime 顶点v的时间标签fTime
	FTime(v int) int
	// Parent 顶点v在遍历树中的父亲节点编号，小于0时不存在父节点
	Parent(v int) int
	// Priority 顶点v在遍历树中的优先级数
	Priority(v int) int

	// EdgeNumbers 边总数
	EdgeNumbers() int
	// EdgeExists 边(v, u)是否存在
	EdgeExists(v, u int) bool
	// EdgeInsert 在顶点v和u之间插入权重为w的边e, 已存在边时替换权重并返回旧权值
	EdgeInsert(v, u int, w interface{}) interface{}
	// EdgeRemove 删除顶点v和u之间的边e，返回该边信息
	EdgeRemove(v, u int) interface{}
	// EdgeType 边(v, u)的类型
	EdgeType(v, u int) EdgeType
	// EdgeWeight 边(v, u)的权重
	EdgeWeight(v, u int, w interface{}) interface{}

	// Reset 所有图结构信息复位
	Reset()

	// BFS 从s开始对全图执行广度有限搜索
	BFS(s int)
	// DFS 从s开始对全图执行深度有限搜索
	DFS(s int)
	// TSort 基于DFS的拓扑排序算法
	TSort(s int) []interface{}
	// Bcc 基于DFS的BCC(双联通域)分解算法
	Bcc(s int)
}
