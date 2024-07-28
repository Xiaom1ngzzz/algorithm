package algo

type MyCircularQueue struct {
	queue []int
	l     int
	r     int
	size  int
	limit int
}

func NewMyCircularQueue(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		l:     0,
		r:     0,
		size:  0,
		limit: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.r] = value
	if this.r == this.limit-1 {
		this.r = 0
	} else {
		this.r++
	}
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.l == this.limit-1 {
		this.l = 0
	} else {
		this.l++
	}
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.l]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	var last int
	if this.r == 0 {
		last = this.limit - 1
	} else {
		last = this.r - 1
	}
	return this.queue[last]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.limit
}
