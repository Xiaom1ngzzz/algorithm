package leetcode2

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

func addTwoNumbers(h1, h2 *ListNode) *ListNode {
	var ans *ListNode = nil
	var cur *ListNode = nil
	carry := 0

	for h1 != nil || h2 != nil {
		sum := carry
		if h1 == nil {
			sum += 0
		} else {
			sum += h1.val
		}
		if h2 == nil {
			sum += 0
		} else {
			sum += h2.val
		}

		val := sum % 10
		carry = sum / 10
		if ans == nil {
			ans = NewListNode(val)
			cur = ans
		} else {
			cur.next = NewListNode(val)
			cur = cur.next
		}

		if h1 != nil {
			h1 = h1.next
		}
		if h2 != nil {
			h2 = h2.next
		}
	}
	if carry == 1 {
		cur.next = NewListNode(1)
	}
	return ans
}
