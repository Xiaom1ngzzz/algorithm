package stack

// stack.go 实现了一个简单的栈数据结构
// 包含Push、Pop、Peek、IsEmpty和Size方法
// 使用interface{}类型可以存储任意类型的数据
// 使用切片实现栈
// 使用错误处理机制处理栈为空的情况
// 使用len函数获取栈的大小
// 使用append函数向栈中添加元素
// 使用切片的切片操作移除栈顶元素
// 使用切片的最后一个元素作为栈顶元素

import (
	"errors"
)

// Stack 定义一个栈结构
type Stack struct {
	elements []interface{} // 使用interface{}类型可以存储任意类型的数据
}

// NewStack 创建一个新的栈
func NewStack() *Stack {
	return &Stack{}
}

// Push 向栈中添加元素
func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

// Pop 从栈中移除并返回栈顶元素
func (s *Stack) Pop() (interface{}, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

// Peek 返回栈顶元素但不移除它
func (s *Stack) Peek() (interface{}, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

// IsEmpty 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size 返回栈中元素的数量
func (s *Stack) Size() int {
	return len(s.elements)
}
