package leetcode260

// 数组中有2种数出现了奇数次，其他的数都出现了偶数次
// 返回这2种出现了奇数次的数

func singleNumberIII(nums []int) []int {
	eor1 := 0
	for _, num := range nums {
		eor1 ^= num
	}
	rightOne := eor1 & ((^int(0) ^ eor1) + 1)
	eor2 := 0
	for _, num := range nums {
		if eor2&rightOne == 0 {
			eor2 ^= num
		}
	}
	return []int{eor2, eor2 ^ eor1}
}
