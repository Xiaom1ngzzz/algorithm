package leetcode9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	offset := 1

	for x/offset >= 10 {
		offset *= 10
	}

	// 首尾判断
	for x != 0 {
		if x/offset != x%10 {
			return false
		}
		x = (x % offset) / 10
		offset /= 100
	}
	return true
}
