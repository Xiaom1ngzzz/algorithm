package leetcode160

import "math"

// 返回两个无环链表相交的第一个节点

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB
	diff := 0 // headA和headB的长度差
	for a.Next != nil {
		a = a.Next
		diff++
	}
	for b.Next != nil {
		b = b.Next
		diff--
	}

	// a为headA的最后一个节点
	// b为headB的最后一个节点
	if a != b { // 连个链表不相交
		return nil
	}

	if diff >= 0 {
		a = headA
		b = headB
	} else {
		a = headB
		b = headA
	}

	diff = int(math.Abs(float64(diff)))

	for diff != 0 {
		a = a.Next
		diff--
	}
	for a != b {
		a = a.Next
		b = b.Next
	}
	return a
}
