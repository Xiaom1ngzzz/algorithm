package leetcode234

type ListNode struct {
	Val  int
	Next *ListNode
}

// isPalindrome 用于判断链表是否是回文结构
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 双指针找中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 现在中点就是slow，从终点开始往后逆序
	pre := slow
	cur := pre.Next
	pre.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// 上面的过程已经把链表调整成从左右两侧往中间指
	// head -> ... -> slow <- ... <- pre
	ans := true
	left := head
	right := pre
	// left往右，right往左，每一步对比值是否相等，如果某一步不一样答案就是false
	for left != nil && right != nil {
		if left.Val != right.Val {
			ans = false
			break
		}
		left = left.Next
		right = right.Next
	}

	// 还原原链表
	cur = pre.Next
	pre.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return ans
}
