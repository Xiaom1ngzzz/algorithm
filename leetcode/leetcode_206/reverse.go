package leetcode206

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

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var next *ListNode = nil

	for head != nil {
		next = head.next // 保存下一个节点的引用
		head.next = pre  // 修改当前节点的指针方向
		pre = head       // 更新 pre 指针
		head = next      // 移动到下一个节点
	}

	return pre // 返回新的头节点
}

// DoubleListNode 双链表节点
type DoubleListNode struct {
	val  int
	next *DoubleListNode
	last *DoubleListNode
}

// NewDoubleListNode 创建新节点
func NewDoubleListNode(val int) *DoubleListNode {
	return &DoubleListNode{
		val:  val,
		next: nil,
		last: nil,
	}
}

func reverseDoubleList(head *DoubleListNode) *DoubleListNode {
	var pre *DoubleListNode = nil
	var next *DoubleListNode = nil

	for head != nil {
		next = head.next
		head.next = pre
		head.last = next
		pre = head
		head = next
	}
	return pre
}

func linkedListInversion(head *ListNode) *ListNode {
	/*
		dummy := &ListNode{} // 使用虚拟头节点

		for head != nil {
			next := head.next

			head.next = dummy
			dummy = head
			head = next
		}

	*/
	var pre *ListNode = nil

	for head != nil {
		next := head.next

		head.next = pre
		pre = head
		head = next
	}

	return pre
}
