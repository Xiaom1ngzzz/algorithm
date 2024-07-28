package leetcode641

// 设计循环双端队列

type MyCircularDeque struct {
	queue []int
	l     int
	r     int
	size  int
	limit int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: make([]int, k),
		size:  0,
		l:     0,
		r:     0,
		limit: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.l, this.r = 0, 0
		this.queue[0] = value
	} else {
		if this.l == 0 {
			this.l = this.limit - 1
		} else {
			this.l--
		}
		this.queue[this.l] = value
	}
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.l, this.r = 0, 0
		this.queue[0] = value
	} else {
		if this.r == this.limit-1 {
			this.r = 0
		} else {
			this.r++
		}
		this.queue[this.r] = value
	}
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
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

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	if this.r == 0 {
		this.r = this.limit - 1
	} else {
		this.r--
	}
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.l]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.queue[this.r]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == this.limit
}
