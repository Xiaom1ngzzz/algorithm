package leetcode162

/*
有序数组中找>=num的最左位置
*/

func rightFindLeft(arr []int, target int) int {
	for index, v := range arr {
		if v >= target {
			return index
		}
	}
	return -1
}

func findLeft(arr []int, target int) int {
	l, r, m := 0, len(arr)-1, 0
	ans := -1
	for l <= r {
		m = l + ((r - l) >> 1)
		if arr[m] >= target {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}
