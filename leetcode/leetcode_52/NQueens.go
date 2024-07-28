package leetcode52

import "math"

// 用数组表示路径实现的N皇后问题
func totalNQueens1(n int) int {
	if n < 1 {
		return 0
	}

	path := make([]int, n)

	return f1(0, &path, n)
}

// i ：当前来到的行
// path ：0...i-1行的皇后，都摆在了那些列
// n ： 问题的规模，是几皇后的问题
func f1(i int, path *[]int, n int) int {
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ {
		if check(path, i, j) {
			(*path)[i] = j
			res += f1(i+1, path, n)
		}
	}
	return res
}

// 当前在i行j列的位置，摆了一个皇后
// 0...i-1行的皇后状况，path[0...i-1]
// 返回会不会冲突，不会冲突，有效！true
// 会冲突，无效，返回false
func check(path *[]int, i, j int) bool {
	for k := 0; k < j; k++ {
		if j == (*path)[k] || math.Abs(float64(i-k)) == math.Abs(float64((*path)[k]-j)) {
			return false
		}
	}
	return true
}

func totalNQueens2(n int) int {
	if n < 1 {
		return 0
	}
	// n = 5
	// 1 << 5 = 0...100000 - 1
	// limit  = 0...011111
	limit := (1 << n) - 1
	return f2(limit, 0, 0, 0)
}

// limit ： 当前是几皇后问题
// col  ： 之前皇后的列影响
// left ： 之前皇后的右上->左下对角线影响
// right： 之前皇后的左上->右下对角线影响
func f2(limit, col, left, right int) int {
	if col == limit {
		// 所有皇后都放完了
		return 1
	}
	// 总限制
	ban := col | left | right
	// ^ban ： 1可放皇后，0不能放
	// 可以放置皇后的范围
	candidate := limit & (^ban)
	place := 0
	ans := 0
	for candidate != 0 {
		place = candidate & (-candidate)
		candidate ^= place
		ans += f2(limit, col|place, (left|place)>>1, (right|place)<<1)
	}
	return ans
}
