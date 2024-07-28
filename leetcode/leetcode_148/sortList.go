package leetcode148

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// sortList 链表排序
func sortList(head *ListNode) *ListNode {
	// 统计链表的长度
	n := 0
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}

	// 使用循环的方式排序
	for step := 1; step < n; step <<= 1 {
		l1 := head
		r1 := findEnd(l1, step)
		l2 := r1.Next
		r2 := findEnd(l2, step)
		next := r2.Next
		r1.Next = nil
		r2.Next = nil
		merge(l1, r1, l2, r2)
		head = start
		lastTeamEnd := end

		for next != nil {
			l1 = next
			r1 = findEnd(l1, step)
			l2 = r1.Next
			if l2 == nil {
				lastTeamEnd.Next = l1
				break
			}
			r2 = findEnd(l2, step)
			next = r2.Next
			r1.Next = nil
			r2.Next = nil
			merge(l1, r1, l2, r2)
			lastTeamEnd.Next = start
			lastTeamEnd = end
		}
	}
	return head
}

// findEnd 查找分组的最后一个节点
func findEnd(s *ListNode, k int) *ListNode {
	fmt.Printf("%v\n", s)
	k--
	for s.Next != nil && k != 0 {
		s = s.Next
		k--
	}
	return s
}

var (
	start *ListNode
	end   *ListNode
)

// merge 合并两组
func merge(l1, r1, l2, r2 *ListNode) {
	var pre *ListNode
	if l1.Val <= l2.Val {
		start = l1
		pre = l1
		l1 = l1.Next
	} else {
		start = l2
		pre = l2
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		pre.Next = l1
		end = r1
	}

	if l2 != nil {
		pre.Next = l2
		end = r2
	}
}
