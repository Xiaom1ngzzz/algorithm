package leetcode21

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	p1 := NewListNode(3)
	p2 := NewListNode(5)
	p3 := NewListNode(9)
	p4 := NewListNode(10)
	p1.next = p2
	p2.next = p3
	p3.next = p4

	p11 := NewListNode(2)
	p22 := NewListNode(7)
	p33 := NewListNode(8)
	p44 := NewListNode(13)
	p11.next = p22
	p22.next = p33
	p33.next = p44

	h := mergeTwoLists(p1, p11)

	for h != nil {
		fmt.Println(h.val)
		h = h.next
	}
}

func TestMergeTwoLists2(t *testing.T) {
	p1 := NewListNode(3)
	p2 := NewListNode(5)
	p3 := NewListNode(9)
	p4 := NewListNode(10)
	p1.next = p2
	p2.next = p3
	p3.next = p4

	p11 := NewListNode(2)
	p22 := NewListNode(7)
	p33 := NewListNode(8)
	p44 := NewListNode(13)
	p11.next = p22
	p22.next = p33
	p33.next = p44

	h := mergeTwoLists2(p1, p11)

	for h != nil {
		fmt.Println(h.val)
		h = h.next
	}
}
