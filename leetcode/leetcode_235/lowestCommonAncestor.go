package leetcode235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// root从上到下
	// 如果先遇到了p，说明p是答案
	// 如果先遇到了q，说明q是答案
	// 如果root在p～q之间，不管p和q谁大谁小，只要root在中间，那么此时的root就是答案
	// 如果root在p～q的值的左侧，那么root往右移动
	// 如果root在p～q的值的右侧，那么root往左移动
	for root.Val != p.Val && root.Val != q.Val {
		//
		if min(p.Val, q.Val) < root.Val && root.Val < max(p.Val, q.Val) {
			break
		}
		if root.Val < min(p.Val, q.Val) {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return root
}
