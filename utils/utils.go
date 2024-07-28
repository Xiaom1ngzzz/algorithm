package utils

import (
	"math/rand"
	"sort"
)

func RandomArray(n, v int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr = append(arr, int(rand.Float64()*float64(v))+1)
	}
	return arr
}

func SameArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	n := len(arr1)
	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func ReverseSlice(slice []int) []int {
	length := len(slice)
	reversed := make([]int, length)

	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[j] = slice[i]
	}

	return reversed
}

func SlicesAreEqual(slice1, slice2 []int) bool {
	// 首先检查切片长度是否相同
	if len(slice1) != len(slice2) {
		return false
	}

	sort.SliceStable(slice1, func(i int, j int) bool {
		return slice1[i] < slice1[j]
	})
	sort.SliceStable(slice2, func(i int, j int) bool {
		return slice2[i] < slice2[j]
	})

	// 逐个比较切片中的元素
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	// 如果所有元素都相同，则切片相同
	return true
}
