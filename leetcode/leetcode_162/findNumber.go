package leetcode162

// 为了验证
func right(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

// 查找有序数组中给定的数是否存在
func exist(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}
	l, r, m := 0, len(arr)-1, 0

	for l <= r {
		m = l + ((r - l) >> 1)
		if arr[m] == target {
			return true
		} else if arr[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
