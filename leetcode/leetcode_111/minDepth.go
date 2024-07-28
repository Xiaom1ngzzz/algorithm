package leetcode111

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	ldeep, rdeep := math.MaxInt, math.MaxInt
	if root.Left != nil {
		ldeep = minDepth(root.Left)
	}
	if root.Right != nil {
		rdeep = minDepth(root.Right)
	}

	return min(ldeep, rdeep) + 1
}
