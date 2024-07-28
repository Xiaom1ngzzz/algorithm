package leetcode136

// 数组中1种数出现了奇数次，其他的数都出现了偶数次
// 返回出现了奇数次的数

func singleNumber(arr []int) int {
	eor := 0
	for _, v := range arr {
		eor ^= v
	}
	return eor
}
