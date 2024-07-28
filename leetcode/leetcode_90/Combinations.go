package leetcode90

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	path := make([]int, len(nums))

	f(nums, 0, &path, 0, &ans)
	return ans
}

func f(nums []int, i int, path *[]int, size int, ans *[][]int) {
	if i == len(nums) {
		cur := make([]int, 0)
		for j := 0; j < size; j++ {
			cur = append(cur, (*path)[j])
		}
		*ans = append(*ans, cur)
	} else {
		// 下一组的第一个数的位置
		j := i + 1
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}
		f(nums, j, path, size, ans)
		// x的数量
		for ; i < j; i++ {
			(*path)[size] = nums[i]
			size++
			f(nums, j, path, size, ans)
		}
	}
}
