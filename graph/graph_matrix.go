package graph

import (
	"data-structures-and-algorithms/contract"
	"data-structures-and-algorithms/queue"
	"data-structures-and-algorithms/stack"
	"math"
)

var _ Graph = (*MatrixGraph)(nil)

// MatrixGraph 以邻接矩阵形式实现的图
type MatrixGraph struct {
	v, e           int       // 节点数，边数
	nodes          []*vertex // 节点集
	edges          [][]*edge // 边集：基于二维切片( len(nodes) * len(nodes) )
	ring           bool      // 是否存在环
	edgeComparator contract.Comparator
}

// NewMatrixGraph 初始化空矩阵图
func NewMatrixGraph(comparators ...contract.Comparator) *MatrixGraph {
	edgeCmp := contract.DefaultComparator
	return &MatrixGraph{edgeComparator: edgeCmp}
}

// NodeNumbers 节点总数
func (g *MatrixGraph) NodeNumbers() int {
	return g.v
}

// Exists 节点 v 是否存在
func (g *MatrixGraph) Exists(v int) bool {
	return v >= 0 && v < g.v
}

// Insert 插入顶点，返回编号
func (g *MatrixGraph) Insert(data interface{}) int {
	g.nodes = append(g.nodes, newVertex(data))
	for i := 0; i < g.v; i++ { // 二维矩阵新增一列
		g.edges[i] = append(g.edges[i], nil)
	}
	g.v++
	// 二维矩阵新增一行(长度为 g.v)
	g.edges = append(g.edges, make([]*edge, g.v))
	return g.v - 1
}

// Remove 删除顶点v及其关联边，返回该顶点信息
func (g *MatrixGraph) Remove(v int) interface{} {
	// 移除出边 (一行)
	for i := 0; i < g.v; i++ {
		if g.EdgeExists(v, i) {
			g.e--
			g.nodes[i].in--
		}
		g.edges[v][i] = nil
	}
	g.v--
	data := g.nodes[v].data
	if v < g.v {
		copy(g.edges[v:], g.edges[v+1:]) // 移除 v 的节点
		copy(g.nodes[v:], g.nodes[v+1:]) // 移除 v 的边行
	}
	g.edges[g.v] = nil
	g.edges = g.edges[:g.v]
	// 移除节点
	g.nodes[g.v] = nil
	g.nodes = g.nodes[:g.v]
	// 移除入边（一列）
	for i := 0; i < g.v; i++ {
		if g.EdgeExists(i, v) {
			g.e--
			g.nodes[i].out--
		}
		if v < g.v { // 每行切片前移
			copy(g.edges[i][v:], g.edges[i][v+1:])
		}
		g.edges[i][g.v] = nil
		g.edges[i] = g.edges[i][:g.v]
	}
	return data
}

// Vertex 顶点v的数据，不存在时返回nil
func (g *MatrixGraph) Vertex(v int) interface{} {
	return g.nodes[v].data
}

// InDegree 顶点v的入度
func (g *MatrixGraph) InDegree(v int) int {
	return g.nodes[v].in
}

// OutDegree 顶点v的出度
func (g *MatrixGraph) OutDegree(v int) int {
	return g.nodes[v].out
}

// FirstNbr 顶点v的首个邻接顶点
func (g *MatrixGraph) FirstNbr(v int) int {
	return g.NextNbr(v, g.v)
}

// NextNbr 顶点v的, 相对于顶点i的下一邻接顶点
func (g *MatrixGraph) NextNbr(v int, i int) int {
	for i = i - 1; -1 < i && !g.EdgeExists(v, i); i-- {
	}
	return i
}

// Status 顶点v的状态
func (g *MatrixGraph) Status(v int) NodeState {
	return g.nodes[v].state
}

// DTime 顶点v的时间标签dTime
func (g *MatrixGraph) DTime(v int) int {
	return g.nodes[v].dTime
}

// FTime 顶点v的时间标签fTime
func (g *MatrixGraph) FTime(v int) int {
	return g.nodes[v].fTime
}

// Parent 顶点v在遍历树中的父亲节点编号，小于0时不存在父节点
func (g *MatrixGraph) Parent(v int) int {
	return g.nodes[v].parent
}

// Priority 顶点v在遍历树中的优先级数
func (g *MatrixGraph) Priority(v int) int {
	return g.nodes[v].priority
}

// EdgeNumbers 边总数
func (g *MatrixGraph) EdgeNumbers() int {
	return g.e
}

// EdgeExists 边(v, u)是否存在
func (g *MatrixGraph) EdgeExists(v, u int) bool {
	return g.edges[v][u] != nil
}

// EdgeInsert 在顶点v和u之间插入权重为w的边e
func (g *MatrixGraph) EdgeInsert(v, u int, w interface{}) (ow interface{}) {
	if g.EdgeExists(v, u) {
		ow = g.edges[v][u].weight
		g.edges[v][u].weight = w
		return
	}
	g.e++
	g.nodes[v].out++
	g.nodes[u].in++
	g.edges[v][u] = newEdge(w)
	return
}

// EdgeRemove 删除顶点v和u之间的边e，返回该边信息
func (g *MatrixGraph) EdgeRemove(v, u int) interface{} {
	if !g.EdgeExists(v, u) {
		return nil
	}
	w := g.edges[v][u].weight
	g.e--
	g.nodes[v].out--
	g.nodes[u].in--
	g.edges[v][u] = nil
	return w
}

// EdgeType 边(v, u)的类型
func (g *MatrixGraph) EdgeType(v, u int) EdgeType {
	return g.edges[v][u].genre
}

// EdgeWeight 边(v, u)的权重
func (g *MatrixGraph) EdgeWeight(v, u int, w interface{}) interface{} {
	return g.edges[v][u].weight
}

// Reset 初始化图
func (g *MatrixGraph) Reset() {
	for i, size := 0, g.v; i < size; i++ {
		g.nodes[i].reset()
		for j := 0; j < size; j++ {
			if g.EdgeExists(i, j) {
				g.edges[i][j].reset()
			}
		}
	}
	g.ring = false
}

// BFS 从s开始对全图执行广度有限搜索
func (g *MatrixGraph) BFS(s int) {
	v, clock := s, 0
	g.Reset()
	g.bfs(v, &clock)
	for v = v + 1; v != s; v = (v + 1) % g.v {
		if g.nodes[v].state == UNDISCOVERED {
			g.bfs(v, &clock)
		}
	}
}

// DFS 从s开始对全图执行深度有限搜索
func (g *MatrixGraph) DFS(s int) {
	v, clock := s, 0
	g.Reset()
	g.dfs(v, &clock, nil, nil)
	for v = v + 1; v != s; v = (v + 1) % g.v {
		if g.nodes[v].state == UNDISCOVERED {
			g.dfs(v, &clock, nil, nil)
		}
	}
}

// 单连通域bfs
func (g *MatrixGraph) bfs(s int, clock *int) {
	que := queue.New()
	que.Push(s)
	for !que.Empty() {
		s = que.Pop().(int)
		*clock++
		g.nodes[s].dTime = *clock
		for v := g.FirstNbr(s); v != -1; v = g.NextNbr(s, v) {
			if g.Status(v) == UNDISCOVERED {
				que.Push(v)
				g.edges[s][v].genre = TREE
				g.nodes[v].parent = s
			} else {
				g.edges[s][v].genre = CROSS
			}
		}
		g.nodes[s].state = VISITED
	}
}

// 单连通域dfs
func (g *MatrixGraph) dfs(s int, clock *int, discover, visited func(*vertex)) {
	*clock++
	g.nodes[s].dTime = *clock
	g.nodes[s].state = DISCOVERED
	if discover != nil {
		discover(g.nodes[s])
	}
	for v := g.FirstNbr(s); v != -1; v = g.NextNbr(s, v) {
		if g.Status(v) == UNDISCOVERED {
			g.edges[s][v].genre = TREE
			g.nodes[v].parent = s
			g.dfs(v, clock, discover, visited)
		} else if g.Status(v) == DISCOVERED {
			g.ring = true
			g.edges[s][v].genre = BACKWARD
		} else {
			if g.nodes[s].dTime < g.nodes[v].dTime {
				g.edges[s][v].genre = FORWARD
			} else {
				g.edges[s][v].genre = CROSS
			}
		}
	}
	if visited != nil {
		visited(g.nodes[s])
	}
	*clock++
	g.nodes[s].fTime = *clock
	g.nodes[s].state = VISITED
}

// TSort 基于DFS的拓扑排序算法：存在环时panic
func (g *MatrixGraph) TSort(s int) []interface{} {
	g.Reset()
	ans := make([]interface{}, 0)
	v, clock := s, 0
	g.Reset()
	visited := func(v *vertex) {
		ans = append(ans, v.data)
	}
	g.dfs(v, &clock, nil, visited)
	for v = v + 1; v != s; v = (v + 1) % g.v {
		if g.nodes[v].state == UNDISCOVERED {
			g.dfs(v, &clock, nil, visited)
		}
	}
	if g.ring {
		panic("graph has ring")
	}
	return ans
}

// Bcc 基于DFS的BCC(双联通域)分解算法
func (g *MatrixGraph) Bcc(s int) {
	g.Reset()
	v, clock := s, 0
	stack := stack.New()
	g.bcc(v, &clock, stack)
	for v = v + 1; v != s; v = (v + 1) % g.v {
		if g.nodes[v].state == UNDISCOVERED {
			g.bcc(v, &clock, stack)
			stack.Pop()
		}
	}
}

func (g *MatrixGraph) bcc(s int, clock *int, stack *stack.Stack) {
	*clock++
	g.nodes[s].dTime = *clock
	g.nodes[s].state = DISCOVERED
	for v := g.FirstNbr(s); v != -1; v = g.NextNbr(s, v) {
		if g.Status(v) == UNDISCOVERED {
			g.edges[s][v].genre = TREE
			g.nodes[v].parent = s
			g.bcc(v, clock, stack)
			if g.nodes[v].fTime < g.nodes[s].dTime {
				g.nodes[s].fTime = int(math.Min(float64(g.nodes[s].fTime), float64(g.nodes[v].fTime)))
			} else {
				for s != stack.Pop().(int) {
				}
				stack.Push(s)
			}
		} else if g.Status(v) == DISCOVERED {
			g.ring = true
			g.edges[s][v].genre = BACKWARD
			if v != g.nodes[s].parent {
				g.nodes[s].fTime = int(math.Min(float64(g.nodes[s].fTime), float64(g.nodes[v].dTime)))
			}
		} else {
			if g.nodes[s].dTime < g.nodes[v].dTime {
				g.edges[s][v].genre = FORWARD
			} else {
				g.edges[s][v].genre = CROSS
			}
		}
	}
	g.nodes[s].state = VISITED
}
