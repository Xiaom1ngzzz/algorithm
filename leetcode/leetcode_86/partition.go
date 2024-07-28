package leetcode86

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

func partition(head *ListNode, x int) *ListNode {
	var leftHead *ListNode = nil
	var leftTail *ListNode = nil
	var rightHead *ListNode = nil
	var rightTail *ListNode = nil
	var next *ListNode = nil

	for head != nil {
		next = head.next
		head.next = nil
		if head.val < x {
			if leftHead == nil {
				leftHead = head
			} else {
				leftTail.next = head
			}
			leftTail = head
		} else {
			if rightHead == nil {
				rightHead = head
			} else {
				rightTail.next = head
			}
			rightTail = head
		}
		head = next
	}
	if leftHead == nil {
		return rightHead
	}
	leftTail.next = rightHead

	return leftHead
}
