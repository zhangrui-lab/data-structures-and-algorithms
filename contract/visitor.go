package contract

// Visitor 值迭代器
type Visitor func(value interface{})

// KvVisitor <key,val>迭代器
type KvVisitor func(key, value interface{})
