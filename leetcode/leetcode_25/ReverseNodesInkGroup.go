package leetcode25

// 每k个节点一组翻转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 输入的头节点为空，直接返回
	if head == nil {
		return head
	}

	start := head
	end := teamEnd(start, k)

	// 输入的链表长度小于 k
	if end == nil {
		return head
	}

	// 输入的链表长度大于 k
	// 特殊处理第一组：因为牵扯到换头的问题
	head = end
	reverse(start, end)
	// 翻转之后start变成了上一组的结尾节点
	lastTeamEnd := start

	for lastTeamEnd.Next != nil {
		start = lastTeamEnd.Next
		end = teamEnd(start, k)
		if end == nil {
			return head
		}
		reverse(start, end)
		lastTeamEnd.Next = end
		lastTeamEnd = start
	}
	return head
}

func teamEnd(s *ListNode, k int) *ListNode {
	k--
	for k != 0 && s != nil {
		s = s.Next
		k--
	}
	return s
}

func reverse(s *ListNode, e *ListNode) {
	e = e.Next
	cur := s
	var pre, next *ListNode
	for cur != e {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	s.Next = e
}
