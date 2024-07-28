package leetcode138

// 复制带随机指针的链表

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
		cur.Next.Next = next
		cur = next
	}

	cur = head
	for cur != nil {
		next := cur.Next.Next
		copy := cur.Next
		if cur.Random != nil {
			copy.Random = cur.Random.Next
		} else {
			copy.Random = nil
		}
		cur = next
	}

	ans := head.Next
	cur = head
	for cur != nil {
		next := cur.Next.Next
		copy := cur.Next
		cur.Next = next
		if next != nil {
			copy.Next = next.Next
		} else {
			copy.Next = nil
		}
		cur = next
	}
	return ans
}
