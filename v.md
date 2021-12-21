* Iterator

```
valid() bool
value() interface{}
set(interface{})
next() Iterator
equal(o Iterator) bool
```

* KvIterator
```
key() interface{}
```

* Bidirectional iterator

```
prev() Iterator
```

* RandomAccessIterator

```
forward(n int) RandomAccessIterator
backward(n int) RandomAccessIterator
at(n int) RandomAccessIterator
```

* Comparator

```
return -1 : a < b
return 0  : a == b
return 1  : a > b
type Comparator func(a, b interface{}) int
```

* IterFunc

```
swap(iter1, iter2 Iterator)
```

* Visitor

```
type ValVisitor func(value interface{}) bool
type KvVisitor func(value interface{}) bool
```

* Seq Container

```
//iterators
begin() Iterator
end() Iterator
// capacity
size() int
empty() int
resize(n int, val interface{})
capacity() int
// element access
front() interface{}
back() interface{}
at() interface{}
// modifiers
assign(n int, val intrface{})
assignRange(first, last Iterator)
insert(i Iterator, value interface{})
insertAt(i int, value interface{})
insertSlice(i Iterator, values[]interface{})
insertAtSlice(n int, values[]interface{})
insertRange(n , first, last Iterator)
insertAtRange(n int, first, last Iterator)
erase(i Iterator)
eraseAt(n int)
pushFront(value interface{})
popFront()
pushBack(value interface{})
popBack()
clear()
swap(o contatiner)
```

* Associative containers

```
//iterators
begin() Iterator
end() Iterator
//element access
at(key interface{}) interface{}
front() interface{}
back() interface{}
// capacity
size() int
empty() int
// modifiers
insert(key, value interface)
get(key) interface
erase(key)
eraseIter(iter)
find(key) iter
clear()
// multi
lowerBound(key interface{}) iter
upperBound(key interface{}) iter
visit(visitor)
```

