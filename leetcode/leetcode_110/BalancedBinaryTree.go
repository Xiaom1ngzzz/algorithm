package leetcode110

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var balance bool

func isBalanced(root *TreeNode) bool {
	balance = true
	height(root)
	return balance
}

func height(cur *TreeNode) int {
	if !balance || cur == nil {
		return 0
	}

	lh := height(cur.Left)
	rh := height(cur.Right)
	if int(math.Abs(float64(lh)-float64(rh))) > 1 {
		balance = false
	}
	return max(lh, rh) + 1
}
