package week02

//LRU Cache的实现

type LRUCache struct {
	capacity int
	size     int
	hMap     map[int]*LinkedNode
	data     *LinkedNode
	head     *LinkedNode
	tail     *LinkedNode
}

// 双向链表数据结构模版
type LinkedNode struct {
	key   int
	value int
	pre   *LinkedNode
	next  *LinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		hMap:     make(map[int]*LinkedNode),
		data:     nil,
		size:     0,
	}
	// 创建保护节点
	head := LinkedNode{
		key:   0,
		value: 0,
	}
	tail := LinkedNode{
		key:   0,
		value: 0,
	}
	head.next = &tail
	tail.pre = &head
	l.data = &head
	l.head = &head
	l.tail = &tail
	return l
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.hMap[key]
	if !ok {
		return -1
	}
	removeNode(val)
	node := this.addToHead(val.key, val.value)
	this.hMap[key] = node
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	val, ok := this.hMap[key]
	this.hMap[key] = this.addToHead(key, value)
	if ok {
		removeNode(val)
		return
	}
	this.size++
	if this.size > this.capacity {
		delete(this.hMap, this.tail.pre.key)
		//超出容量删除尾端数据
		removeNode(this.tail.pre)
	}
}

// 双向链表删除一个节点模版
func removeNode(node *LinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addToHead(key, val int) *LinkedNode {
	node := LinkedNode{
		key:   key,
		value: val,
	}
	this.data.next.pre = &node
	node.pre = this.data
	node.next = this.data.next
	this.data.next = &node
	return &node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
