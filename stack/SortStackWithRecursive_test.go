package stack

import (
	"fmt"
	"testing"
)

func TestSortStackWithRecursive(t *testing.T) {
	stack := NewStack()
	stack.Push(6)
	stack.Push(2)
	stack.Push(5)
	stack.Push(4)
	stack.Push(5)
	stack.Push(1)
	stack.Push(3)

	// Print：3, 1, 5，4，5，2，6

	sort(stack)

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}
