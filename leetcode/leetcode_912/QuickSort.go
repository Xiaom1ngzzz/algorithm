package leetcode912

import (
	"math/rand"
)

var counter int = 0

func quickSort1(arr []int, l, r int) {
	if l >= r {
		return
	}
	counter++

	if counter > 15 {
		return
	}
	// 随机一个索引，范围 l....r之间
	x := arr[l+rand.Intn(r-l+1)]
	mid := partition1(arr, l, r, x)

	quickSort1(arr, l, mid-1)
	quickSort1(arr, mid+1, r)
}

func partition1(arr []int, l, r, x int) int {
	a := l
	xi := 0
	for i := l; i <= r; i++ {
		if arr[i] <= x {
			arr[a], arr[i] = arr[i], arr[a]
			if arr[a] == x {
				xi = a
			}
			a++
		}
	}
	arr[xi], arr[a-1] = arr[a-1], arr[xi]
	return a - 1
}

var (
	first int
	last  int
)

func quickSort2(arr []int, l, r int) {
	if l >= r {
		return
	}

	// 随机选择基准值
	x := arr[l+rand.Intn(r-l+1)]
	partition2(arr, l, r, x)
	left, right := first, last
	quickSort2(arr, l, left-1)
	quickSort2(arr, right+1, r)
}

func partition2(arr []int, l, r, x int) {
	first = l
	last = r
	i := l

	for i <= last {
		if arr[i] == x {
			i++
		} else if arr[i] < x {
			swap(arr, first, i)
			first++
			i++
		} else {
			swap(arr, i, last)
			last--
		}
	}
}
