package leetcode225

import "container/list"

type stack1 struct {
	data list.List
}

func NewStack1() *stack1 {
	return &stack1{
		data: *list.New(),
	}
}

func (s *stack1) isEmpty() bool {
	return s.data.Len() == 0
}

func (s *stack1) push(val int) {
	s.data.PushBack(val)
}

func (s *stack1) pop() int {
	val := s.data.Remove(s.data.Back())
	return val.(int)
}

func (s *stack1) peek() int {
	return s.data.Back().Value.(int)
}

func (s *stack1) size() int {
	return s.data.Len()
}

type stack2 struct {
	data      []int
	stackSize int
}

func NewStack2(n int) *stack2 {
	return &stack2{
		data:      make([]int, n),
		stackSize: 0,
	}
}

func (s *stack2) isEmpty() bool {
	return s.stackSize == 0
}

func (s *stack2) push(val int) {
	s.data[s.stackSize] = val
	s.stackSize++
}

func (s *stack2) pop() int {
	if s.isEmpty() {
		return -1
	}
	s.stackSize--
	return s.data[s.stackSize]
}

func (s *stack2) peek() int {
	return s.data[s.stackSize-1]
}

func (s *stack2) size() int {
	return s.stackSize
}
