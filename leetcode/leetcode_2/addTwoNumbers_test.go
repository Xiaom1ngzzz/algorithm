package leetcode2

import (
	"fmt"
	"testing"

	"github.com/Ljazz/algorithm/utils"
)

func TestAddTwoNumbers(t *testing.T) {
	s1 := []int{1, 0, 7, 4, 6}
	p1 := NewListNode(4)
	p2 := NewListNode(8)
	p3 := NewListNode(9)
	p1.next = p2
	p2.next = p3

	p11 := NewListNode(2)
	p22 := NewListNode(6)
	p33 := NewListNode(7)
	p44 := NewListNode(9)
	p11.next = p22
	p22.next = p33
	p33.next = p44

	h := addTwoNumbers(p1, p11)
	var res []int
	for h != nil {
		// fmt.Println(h.val)
		res = append(res, h.val)
		h = h.next
	}
	if !utils.SlicesAreEqual(s1, res) {
		fmt.Println("出错了！！！")
	}
}
