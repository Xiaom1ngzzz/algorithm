package leetcode21

// ListNode 链表节点
type ListNode struct {
	val  int
	next *ListNode
}

// NewListNode 创建新节点
func NewListNode(val int) *ListNode {
	return &ListNode{
		val:  val,
		next: nil,
	}
}

func mergeTwoLists(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var head *ListNode
	var cur1 *ListNode
	var cur2 *ListNode
	var pre *ListNode

	if head1.val <= head2.val {
		head = head1
	} else {
		head = head2
	}
	cur1 = head.next
	if head == head1 {
		cur2 = head2
	} else {
		cur2 = head1
	}
	pre = head

	for cur1 != nil && cur2 != nil {
		if cur1.val <= cur2.val {
			pre.next = cur1
			cur1 = cur1.next
		} else {
			pre.next = cur2
			cur2 = cur2.next
		}
		pre = pre.next
	}
	if cur1 != nil {
		pre.next = cur1
	} else {
		pre.next = cur2
	}
	return head
}

func mergeTwoLists2(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{} // 使用虚拟头节点
	cur := dummy         // 当前节点

	for head1 != nil && head2 != nil {
		if head1.val <= head2.val {
			cur.next = head1
			head1 = head1.next
		} else {
			cur.next = head2
			head2 = head2.next
		}
		cur = cur.next
	}

	// 将剩余部分直接接到合并链表的尾部
	if head1 != nil {
		cur.next = head1
	} else {
		cur.next = head2
	}

	return dummy.next // 返回虚拟头节点的下一个节点，即合并后的头节点
}
