package leetcode98

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MAXN = 10001

var (
	stack [MAXN]*TreeNode
	r     int
)

func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var pre *TreeNode
	r = 0
	for r > 0 || root != nil {
		if root != nil {
			stack[r] = root
			r++
			root = root.Left
		} else {
			r--
			root = stack[r]
			if pre != nil && pre.Val >= root.Val {
				return false
			}
			pre = root
			root = root.Right
		}
	}
	return true
}

var (
	tMax int
	tMin int
)

func isValidBST(root *TreeNode) bool {
	if root == nil {
		tMin = math.MaxInt
		tMax = math.MinInt
		return true
	}
	lok := isValidBST(root.Left)
	lmin := tMin
	lmax := tMax
	rok := isValidBST(root.Right)
	rmin := tMin
	rmax := tMax
	// 整棵树的最小值
	tMin = min(min(lmin, rmin), root.Val)
	// 整棵树的最大值
	tMax = max(max(lmax, rmax), root.Val)
	return lok && rok && lmax < root.Val && root.Val < rmin
}
