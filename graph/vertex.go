package graph

import "math"

// NodeState 顶点状态类型
type NodeState uint8

const (
	_            NodeState = iota
	UNDISCOVERED           //  未开始访问
	DISCOVERED             // 已开始访问，暂未访问完成
	VISITED                // 已访问完成
)

// 顶点类型
type vertex struct {
	data         interface{} //数据
	in, out      int         // 出入度
	state        NodeState   // 节点状态
	dTime, fTime int         // 开始，结束访问时间
	parent       int         // 访问时的父节点
	priority     int         // 访问优先级(优先级数值越小时优先级越高)
}

// 新建以data为值的节点
func newVertex(data interface{}) *vertex {
	v := &vertex{data: data}
	v.reset()
	return v
}

// 节点 v 辅助信息复位
func (v *vertex) reset() {
	v.dTime = 0
	v.fTime = 0
	v.parent = -1
	v.state = UNDISCOVERED
	v.priority = math.MaxInt
}
