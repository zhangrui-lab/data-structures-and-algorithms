package bitmap

import "bytes"

// BitMap 位图
type BitMap struct {
	data []byte
}

// NewBitMap 创建一个n位的位图
func NewBitMap(n int) *BitMap {
	b := &BitMap{}
	b.data = make([]byte, (n+7)>>3)
	return b
}

// FromByteSlice 从字节切片构建位图
func FromByteSlice(data []byte) *BitMap {
	b := NewBitMap(len(data) << 3)
	copy(b.data, data)
	return b
}

// Test 测试第i位是否为1
func (b *BitMap) Test(i int) bool {
	b.expand(i)
	return b.data[i>>3]&(0x80>>(i&0x07)) != 0
}

// Set 设置第i位为1
func (b *BitMap) Set(i int) {
	b.expand(i)
	b.data[i>>3] |= 0x80 >> (i & 0x07)
}

// Clear 清除第i位为0
func (b *BitMap) Clear(i int) {
	b.expand(i)
	b.data[i>>3] &= ^(0x80 >> (i & 0x07))
}

// Output 位图中的前n串信息做01串进行输出
func (b *BitMap) Output(n int) string {
	b.expand(n - 1)
	var buff bytes.Buffer
	for i := 0; i < n; i++ {
		if b.Test(i) {
			buff.WriteString("1")
		} else {
			buff.WriteString("0")
		}
	}
	return buff.String()
}

// Clone 输出前n位的位模式
func (b *BitMap) Clone(n int) *BitMap {
	b.expand(n)
	o := NewBitMap(n)
	copy(o.data, b.data)
	return o
}

// 将位图扩展到包含第i位
func (b *BitMap) expand(i int) {
	for i >= cap(b.data)*8 {
		b.data = append(b.data, byte(0)) // 使用切片自身的扩容方式
	}
	b.data = b.data[:cap(b.data)]
}
