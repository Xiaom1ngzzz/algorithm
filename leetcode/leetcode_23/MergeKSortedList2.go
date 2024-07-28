package cfde9e5b9b4cf2b6bafa5f3ef33fa6

// 合并k个有序链表

// SinglyListNode Definition for singly-linked list.
type SinglyListNode struct {
	Val  int
	Next *SinglyListNode
}

// var min_heap []*SinglyListNode

func mergeKLists1(lists []*SinglyListNode) *SinglyListNode {
	size := len(lists)
	if size == 0 {
		return nil
	}

	// 建小根堆
	h := make([]*SinglyListNode, size)
	for i := 0; i < size; i++ {
		if lists[i] != nil {
			h[i] = lists[i]
		}
		heapSiftUp(h, i)
	}

	head := h[0]
	pre := head
	if pre.Next != nil {
		h[0] = pre.Next
		heapSiftDown(h, 0, size)
	}

	for size > 0 {
		cur := h[0]
		pre.Next = cur
		pre = cur
		if cur.Next != nil {
			h[0] = cur.Next
		} else {
			size--
			h[0] = h[size]
			heapSiftDown(h, 0, size)
		}
	}
	return head
}

// heapSiftUp 从节点 i 开始，从底至顶堆化
func heapSiftUp(heap []*SinglyListNode, i int) {
	for heap[i] != nil && heap[i].Val < heap[(i-1)/2].Val {
		heap[i], heap[(i-1)/2] = heap[(i-1)/2], heap[i]
		i = (i - 1) / 2
	}
}

// heapSiftDown 从节点 i 开始，从顶至底堆化
func heapSiftDown(heap []*SinglyListNode, i, size int) {
	left := i*2 + 1
	right := i*2 + 2
	smallest := i

	if left < size && heap[left] != nil && heap[left].Val < heap[smallest].Val {
		smallest = left
	}
	if right < size && heap[right] != nil && heap[right].Val < heap[smallest].Val {
		smallest = right
	}

	if smallest != i {
		heap[i], heap[smallest] = heap[smallest], heap[i]
		heapSiftDown(heap, smallest, size)
	}
}
