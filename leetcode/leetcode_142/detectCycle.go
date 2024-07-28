package leetcode142

type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle 用于检测链表是否有环，并返回入环节点
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow := head.Next
	fast := head.Next.Next

	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
