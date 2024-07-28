package leetcode268

// 找到缺失的数字

func missingNumber(arr []int) int {
	eor := 0
	for idx, v := range arr {
		eor = eor ^ idx ^ v
	}
	eor ^= len(arr)
	return eor
}
