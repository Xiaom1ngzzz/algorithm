package leetcode232

import (
	"container/list"
)

// 用链表实现
type Queue1 struct {
	data *list.List
}

func NewQueue1() *Queue1 {
	return &Queue1{
		data: list.New(),
	}
}

func (q *Queue1) isEmpty() bool {
	return q.data.Len() == 0
}

// 添加到尾部
func (q *Queue1) offer(x int) {
	q.data.PushBack(x)
}

// 头部取
func (q *Queue1) poll() int {
	val := q.data.Remove(q.data.Front())
	return val.(int)
}

// 返回队列头的元素，但是不弹出
func (q *Queue1) peek() int {
	return q.data.Front().Value.(int)
}

func (q *Queue1) size() int {
	return q.data.Len()
}

// 用数组实现
type Queue2 struct {
	queue []int
	l     int
	r     int
}

func NewQueue2(n int) *Queue2 {
	return &Queue2{
		queue: make([]int, n),
		l:     0,
		r:     0,
	}
}

func (q *Queue2) isEmpty() bool {
	return q.l == q.r
}

// 添加到尾部
func (q *Queue2) offer(x int) {
	q.queue[q.r] = x
	q.r++
}

// 头部取
func (q *Queue2) poll() int {
	val := q.queue[q.r-1]
	q.queue = q.queue[:q.r-1]
	return val
}

// 返回队列头的元素，但是不弹出
func (q *Queue2) peek() int {
	return q.queue[q.r-1]
}

func (q *Queue2) size() int {
	return q.r - q.l
}
