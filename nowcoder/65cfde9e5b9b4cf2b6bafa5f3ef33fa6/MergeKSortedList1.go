package cfde9e5b9b4cf2b6bafa5f3ef33fa6

import (
	"container/heap"
)

// 合并k个有序链表

type ListNode struct {
	val  int
	next *ListNode
}

type minHeap []*ListNode

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i].val < (*h)[j].val }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() any {
	idx := len(*h) - 1
	last := (*h)[idx]
	*h = (*h)[:idx]
	return last
}

func (h *minHeap) Top() any {
	return (*h)[0]
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &minHeap{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}
	if h.Len() == 0 {
		return nil
	}

	head := heap.Pop(h).(*ListNode)
	pre := head
	if pre.next != nil {
		heap.Push(h, pre.next)
	}
	for h.Len() != 0 {
		cur := heap.Pop(h).(*ListNode)
		pre.next = cur
		pre = cur
		if cur.next != nil {
			heap.Push(h, cur.next)
		}
	}
	return head
}
