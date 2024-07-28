package leetcode912

import "fmt"

var mergeHelp []int = make([]int, 50001)

// 递归版
func mergeSort1(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := l + (r-l)/2
	fmt.Println("l = ", l, "m = ", m, "r = ", r, "len(arr) = ", len(arr))
	mergeSort1(arr, l, m)
	mergeSort1(arr, m+1, r)
	merge(arr, l, m, r)
}

// 非递归版
//
//	arr  3  8  7  6  4  5  1  2
//
// step = 1
//
//	l =                        0  2  4  6  8
//	m = l + step - 1           0  2  4  6  8
//	r = min(l+(step<<1)-1, 7)  1  3  5  7
//
// step = 2
//
//	l =                        0  4  8
//	m = l + step - 1           1  5  9
//	r = min(l+(step<<1)-1, 7)  3  7
//
// step = 4
//
//	l =                        0  8
//	m = l + step - 1           3  11
//	r = min(l+(step<<1)-1, 7)  7
//
// step = 8
//
//	l =                        0  8
//	m = l + step - 1           7  15
//	r = min(l+(step<<1)-1, 7)  7

func mergeSort(arr []int) []int {
	n := len(arr)
	mergeHelp = make([]int, n)

	for step := 1; step < n; step <<= 1 {
		l := 0
		for l < n {
			m := l + step - 1
			if m+1 >= n {
				break
			}
			r := l + (step << 1) - 1
			if r > n-1 {
				r = n - 1
			}
			merge(arr, l, m, r)
			l = r + 1
		}
	}
	return arr
}

func merge(arr []int, l, m, r int) {
	a := l
	b := m + 1
	i := l
	for a <= m && b <= r {
		if arr[a] < arr[b] {
			mergeHelp[i] = arr[a]
			a++
		} else {
			mergeHelp[i] = arr[b]
			b++
		}
		i++
	}
	for a <= m {
		mergeHelp[i] = arr[a]
		a++
		i++
	}
	for b <= r {
		mergeHelp[i] = arr[b]
		b++
		i++
	}
	for j := l; j <= r; j++ {
		arr[j] = mergeHelp[j]
	}
}
