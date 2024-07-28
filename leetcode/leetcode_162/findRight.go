package leetcode162

/*
有序数组中找<=num的最右位置
*/

func rightFindRight(arr []int, target int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] <= target {
			return i
		}
	}
	return -1
}

func findRight(arr []int, target int) int {
	l, r, m := 0, len(arr)-1, 0
	ans := -1
	for l <= r {
		m = l + ((r - l) >> 1)
		if arr[m] <= target {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}
