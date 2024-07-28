package leetcode201

// 数字范围按位与

func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right -= right & (-right)
	}
	return right
}
