package leetcode155

// 最小栈

type MinStack1 struct {
	data []int
	min  []int
	size int
}

func NewMinStack() MinStack1 {
	return MinStack1{
		data: make([]int, 0, 8001),
		min:  make([]int, 0, 8001),
		size: 0,
	}
}

func (this *MinStack1) Push(val int) {
	this.data = append(this.data, val)
	if this.size == 0 || val < this.min[this.size-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[this.size-1])
	}
	this.size++
}

func (this *MinStack1) Pop() {
	this.size--
	this.data = this.data[:this.size]
	this.min = this.min[:this.size]
}

func (this *MinStack1) Top() int {
	return this.data[this.size-1]
}

func (this *MinStack1) GetMin() int {
	return this.min[this.size-1]
}

// 数组实现

const MAXN = 8001

type MinStack struct {
	data [MAXN]int
	min  [MAXN]int
	size int
}

func NewConstructor() MinStack {
	return MinStack{
		size: 0,
	}
}

func (this *MinStack) Push(val int) {
	this.data[this.size] = val

	if this.size == 0 || val < this.min[this.size-1] {
		this.min[this.size] = val
	} else {
		this.min[this.size] = this.min[this.size-1]
	}
	this.size++
}

func (this *MinStack) Pop() {
	this.size--
}

func (this *MinStack) Top() int {
	return this.data[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.size-1]
}
