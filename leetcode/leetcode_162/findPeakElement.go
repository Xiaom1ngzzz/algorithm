package leetcode162

func rightFindPeakElement(arr []int) []int {
	ans := make([]int, len(arr))
	// 只有一个元素直接返回
	n := len(arr)
	if n == 1 {
		ans = append(ans, 0)
	}

	// 数组长度 >= 2
	// 单独验证 0 位置是不是峰值点
	if arr[0] > arr[1] {
		ans = append(ans, 0)
	}

	if arr[n-1] > arr[n-2] {
		ans = append(ans, n-1)
	}

	for i := 1; i < n-1; i++ {
		if arr[i] > arr[i+1] && arr[i] > arr[i-1] {
			ans = append(ans, i)
		}
	}
	return ans
}

func findPeakElement(arr []int) int {
	// 只有一个元素直接返回
	n := len(arr)
	if n == 1 {
		return 0
	}

	// 数组长度 >= 2
	// 单独验证 0 位置是不是峰值点
	if arr[0] > arr[1] {
		return 0
	}

	// 单独验证 n-1 位置是不是峰值点
	if arr[n-1] > arr[n-2] {
		return n - 1
	}

	//
	l, r, m := 1, n-2, 0
	ans := -1

	for l <= r {
		m = l + ((r - l) >> 1)
		if arr[m-1] > arr[m] {
			r = m - 1
		} else if arr[m] < arr[m+1] {
			l = m + 1
		} else {
			ans = m
			break
		}
	}
	return ans
}
