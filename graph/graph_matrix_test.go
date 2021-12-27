package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrixGraph(t *testing.T) {
	// 7 个节点 7 条边的有向图
	n0, n1, n2, n3, n4, n5, n6 := 0, 1, 2, 3, 4, 5, 6
	w02, w06, w23, w34, w65, w62, w60 := 10, 20, 30, 40, 50, 60, 70

	graph := NewMatrixGraph()
	n := graph.Insert(n0)
	assert.Equal(t, n, 0, "graph.Insert(n0) != 0")
	n = graph.Insert(n1)
	assert.Equal(t, n, 1, "graph.Insert(n1) != 1")
	n = graph.Insert(n2)
	assert.Equal(t, n, 2, "graph.Insert(n2) != 2")

	b := graph.Exists(2)
	assert.True(t, b, "graph.Exists(2) != true")
	b = graph.Exists(3)
	assert.False(t, b, "graph.Exists(3) != false")

	b = graph.EdgeExists(0, 2)
	assert.False(t, b, "graph.EdgeExists(0, 2) != false")
	w := graph.EdgeInsert(0, 2, w02)
	assert.Nil(t, w, "graph.EdgeInsert(0, 2, w02) != nil")
	b = graph.EdgeExists(0, 2)
	assert.True(t, b, "graph.EdgeExists(0, 2) != true")

	assert.Equal(t, graph.NodeNumbers(), 3, "graph.NodeNumbers() != 3")
	assert.Equal(t, graph.EdgeNumbers(), 1, "graph.EdgeNumbers() != 1")

	graph.Insert(n3)
	graph.Insert(n4)
	graph.Insert(n5)
	graph.Insert(n6)

	assert.Equal(t, graph.NodeNumbers(), 7, "graph.NodeNumbers() != 7")

	assert.Equal(t, graph.InDegree(0), 0, "graph.InDegree(0) != 0")
	assert.Equal(t, graph.OutDegree(0), 1, "graph.OutDegree(0) != 1")
	assert.Equal(t, graph.InDegree(2), 1, "graph.InDegree(2) != 1")
	assert.Equal(t, graph.OutDegree(2), 0, "graph.OutDegree(2) != 0")

	graph.EdgeInsert(2, 3, w23)
	assert.Equal(t, graph.OutDegree(2), 1, "graph.OutDegree(2) != 1")
	assert.Equal(t, graph.InDegree(3), 1, "graph.InDegree(3) != 1")

	graph.EdgeInsert(0, 6, w06)
	assert.Equal(t, graph.OutDegree(0), 2, "graph.OutDegree(0) != 2")
	assert.Equal(t, graph.FirstNbr(0), 6, "graph.FirstNbr(0) != 6")
	assert.Equal(t, graph.NextNbr(0, graph.FirstNbr(0)), 2, "graph.NextNbr(0, graph.FirstNbr(0)) != 2")
	assert.Equal(t, graph.NextNbr(0, 2), -1, "graph.NextNbr(0, 2) != -1")

	graph.EdgeInsert(3, 4, w34)
	graph.EdgeInsert(6, 5, w65)
	graph.EdgeInsert(6, 2, w62)
	graph.EdgeInsert(6, 0, w60)

	assert.Equal(t, graph.EdgeNumbers(), 7, "graph.EdgeNumbers() != 1")

	assert.True(t, graph.EdgeExists(2, 3), "graph.EdgeExists(2, 3) != true")
	assert.False(t, graph.EdgeExists(3, 2), "graph.EdgeExists(3, 2) != false")

	v := graph.Remove(0) // todo 对节点的删除会影响后续节点的编号
	assert.Equal(t, v.(int), n0, "graph.Remove(0) != n0")
	assert.Equal(t, graph.EdgeNumbers(), 4, "graph.EdgeNumbers() != 4")

	w = graph.EdgeRemove(5, 1)
	assert.Equal(t, w.(int), w62, "graph.EdgeRemove(5, 1) != w62")

}
