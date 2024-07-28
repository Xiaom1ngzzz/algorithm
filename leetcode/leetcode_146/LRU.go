package leetcode146

type DoubleNode struct {
	key   int
	value int
	last  *DoubleNode
	next  *DoubleNode
}

func newDoubleNode(k, v int) *DoubleNode {
	return &DoubleNode{
		key:   k,
		value: v,
		last:  nil,
		next:  nil,
	}
}

type DoubleList struct {
	head *DoubleNode
	tail *DoubleNode
}

func newDoubleList() *DoubleList {
	return &DoubleList{
		head: nil,
		tail: nil,
	}
}

func (d *DoubleList) addNode(newNode *DoubleNode) {
	if newNode == nil {
		return
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.next = newNode
		newNode.last = d.tail
		d.tail = newNode
	}
}

func (d *DoubleList) moveNodeToTail(node *DoubleNode) {
	if d.tail == node {
		return
	}

	if d.head == node {
		d.head = node.next
		d.head.last = nil
	} else {
		// node 前一个节点 next 指向 node 的下一个节点
		node.last.next = node.next
		// node 后一个节点 last 指向 node 的前一个节点
		node.next.last = node.last
	}
	node.last = d.tail
	node.next = nil
	d.tail.next = node
	d.tail = node
}

func (d *DoubleList) removeHead() *DoubleNode {
	if d.head == nil {
		return nil
	}

	ans := d.head
	if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.head = ans.next
		d.head.last = nil
		ans.next = nil
	}
	return ans
}

type LRUCache struct {
	keyNodeMap map[int](*DoubleNode)
	nodeList   *DoubleList
	capacity   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keyNodeMap: make(map[int](*DoubleNode)),
		nodeList:   newDoubleList(),
		capacity:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.keyNodeMap[key]

	if !ok {
		return -1
	}
	this.nodeList.moveNodeToTail(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.keyNodeMap[key]

	if ok {
		node.value = value
		this.nodeList.moveNodeToTail(node)
	} else {
		if this.capacity == len(this.keyNodeMap) {
			oldNode := this.nodeList.removeHead()
			delete(this.keyNodeMap, oldNode.key)
		}
		newNode := newDoubleNode(key, value)
		this.keyNodeMap[key] = newNode
		this.nodeList.addNode(newNode)
	}
}
