package stack

import (
	"fmt"
	"testing"
)

func TestReverseStackWithRecursive(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	reverse(stack)
	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}
