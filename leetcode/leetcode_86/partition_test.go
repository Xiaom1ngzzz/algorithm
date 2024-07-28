package leetcode86

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	p1 := NewListNode(1)
	p2 := NewListNode(4)
	p3 := NewListNode(3)
	p4 := NewListNode(2)
	p5 := NewListNode(5)
	p6 := NewListNode(2)
	p1.next = p2
	p2.next = p3
	p3.next = p4
	p4.next = p5
	p5.next = p6

	// for p1 != nil {
	// 	fmt.Println(p1.val)
	// 	p1 = p1.next
	// }

	h := partition(p1, 3)

	for h != nil {
		fmt.Println(h.val)
		h = h.next
	}
}
