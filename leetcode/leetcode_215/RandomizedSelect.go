package algo

import (
	"math/rand"
)

// 无序数组中第K大的元素

var (
	findKthLargestFirst int
	findKthLargestLast  int
)

func findKthLargest(nums []int, k int) int {
	return randomizedSelect(nums, len(nums)-k) // 转换成第n-k小
}

func randomizedSelect(arr []int, i int) int {
	ans := 0

	// arr  [l......r]
	// arr  [l...findKthLargestFirst-1]  [findKthLargestFirst...findKthLargestLast] [findKthLargestLast+1...r]
	// 判断 i 处于那部分
	for l, r := 0, len(arr)-1; l <= r; {
		findKthLargestPartition(arr, l, r, arr[l+rand.Intn(r-l+1)])
		if i < findKthLargestFirst {
			r = findKthLargestFirst - 1
		} else if i > findKthLargestLast {
			l = findKthLargestLast + 1
		} else {
			ans = arr[i]
			break
		}
	}
	return ans
}

func findKthLargestPartition(arr []int, l, r, x int) {
	findKthLargestFirst = l
	findKthLargestLast = r

	for i := l; i <= findKthLargestLast; {
		if arr[i] == x {
			i++
		} else if arr[i] < x {
			swap(arr, findKthLargestFirst, i)
			findKthLargestFirst++
			i++
		} else {
			swap(arr, i, findKthLargestLast)
			findKthLargestLast--
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
