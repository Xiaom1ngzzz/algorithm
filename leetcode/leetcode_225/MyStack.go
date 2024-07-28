package leetcode225

import "container/list"

type MyStack struct {
	queue *list.List
}

func NewMyStack() *MyStack {
	return &MyStack{
		queue: list.New(),
	}
}

func (s *MyStack) Push(x int) {
	n := s.queue.Len()  // 长度
	s.queue.PushBack(x) // 尾部添加新元素

	for i := 0; i < n; i++ {
		s.queue.PushBack(s.queue.Back())
	}
}

func (s *MyStack) pop() int {
	val := s.queue.Remove(s.queue.Front())
	return val.(int)
}

func (s *MyStack) top() int {
	return s.queue.Back().Value.(int)
}

func (s *MyStack) Empty() bool {
	return s.queue.Len() == 0
}
