package leetcode232

import (
	"container/list"
)

type MyQueue struct {
	in  *list.List
	out *list.List
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		in:  list.New(),
		out: list.New(),
	}
}

func (q *MyQueue) inToOut() {
	if q.out.Len() == 0 {
		for q.in.Len() > 0 {
			val := q.in.Remove(q.in.Back())
			q.out.PushBack(val)
		}
	}
}

func (q *MyQueue) Push(x int) {
	q.in.PushBack(x)
	q.inToOut()
}

func (q *MyQueue) Pop() int {
	q.inToOut()
	val := q.out.Remove(q.out.Back())
	return val.(int)
}

func (q *MyQueue) Peek() int {
	q.inToOut()
	val := q.out.Back().Value
	return val.(int)
}

func (q *MyQueue) Empty() bool {
	return q.in.Len() == 0 && q.out.Len() == 0
}
