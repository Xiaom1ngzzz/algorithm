package leetcode46

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	f(nums, 0, &ans)
	return ans
}

func f(nums []int, i int, ans *[][]int) {
	if i == len(nums) {
		cur := make([]int, 0)
		for _, v := range nums {
			cur = append(cur, v)
		}
		*ans = append(*ans, cur)
	} else {
		for j := i; j < len(nums); j++ {
			swap(nums, i, j)
			f(nums, i+1, ans)
			swap(nums, i, j) // 回溯 恢复到原来状态
		}
	}
}

func swap(nums []int, i, j int) { nums[i], nums[j] = nums[j], nums[i] }
