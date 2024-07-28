package leetcode2208

import "container/heap"

// 将数组和减半的最少操作次数

func halveArray1(nums []int) int {
	h := &minHeap{}
	heap.Init(h)
	var sum float64 = 0
	for i := 0; i < len(nums); i++ {
		heap.Push(h, float64(nums[i]))
		sum += float64(nums[i])
	}

	var halfSum float64 = 0
	counter := 0

	for halfSum < sum/2 {
		tmp := heap.Pop(h).(float64)
		halfSum += (tmp / 2)
		counter++
		heap.Push(h, tmp/2)
	}
	return counter
}

type minHeap []float64

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Push(x any)         { *h = append(*h, x.(float64)) }
func (h *minHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}
