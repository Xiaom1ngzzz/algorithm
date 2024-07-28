package leetcode295

import (
	"container/heap"
	"math"
)

type minHeap []int

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}
func (h *minHeap) Top() int           { return (*h)[0] }
func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

type maxHeap []int

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}
func (h *maxHeap) Top() int           { return (*h)[0] }
func (h *maxHeap) Len() int           { return len(*h) }
func (h *maxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *maxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

type MedianFinder struct {
	MaxHeap *maxHeap
	MinHeap *minHeap
}

func Constructor() MedianFinder {
	max_h := &maxHeap{}
	min_h := &minHeap{}

	heap.Init(max_h)
	heap.Init(min_h)

	return MedianFinder{
		MaxHeap: max_h,
		MinHeap: min_h,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.MaxHeap.Len() == 0 || this.MaxHeap.Top() >= num {
		heap.Push(this.MaxHeap, num)
	} else {
		heap.Push(this.MinHeap, num)
	}
	this.balance()
}

func (this *MedianFinder) FindMedian() float64 {
	if this.MaxHeap.Len() == this.MinHeap.Len() {
		return float64(this.MaxHeap.Top()+this.MinHeap.Top()) / 2.0
	} else if this.MaxHeap.Len() > this.MinHeap.Len() {
		return float64(this.MaxHeap.Top())
	} else {
		return float64(this.MinHeap.Top())
	}
}

func (this *MedianFinder) balance() {
	if int(math.Abs(float64(this.MaxHeap.Len())-float64(this.MinHeap.Len()))) == 2 {
		if this.MaxHeap.Len() > this.MinHeap.Len() {
			heap.Push(this.MinHeap, heap.Pop(this.MaxHeap))
		} else {
			heap.Push(this.MaxHeap, heap.Pop(this.MinHeap))
		}
	}
}
