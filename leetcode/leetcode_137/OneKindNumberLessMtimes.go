package leetcode137

// 数组中只有1种数出现次数少于m次，其他数都出现了m次
// 返回出现次数小于m次的那种数

func oneKindNumberLessMtimes(nums []int, m int32) int {
	var cnts [32]int32
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			cnts[i] += (int32(num) >> i) & 1
		}
	}
	ans := int32(0)
	for i := 31; i >= 0; i-- {
		if cnts[i]%m != 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
