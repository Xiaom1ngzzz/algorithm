package leetcode958

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MAXN = 101

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue [MAXN]*TreeNode
	l, r := 0, 0

	queue[r] = root
	r++
	// 是否遇到过左右两个孩子不双全的节点
	leaf := false
	for l < r {
		cur := queue[l]
		l++
		// 1. 有右无左
		// 2. 一旦发现孩子不全的节点，接下来必须全是叶节点
		if (cur.Left == nil && cur.Right != nil) || (leaf && (cur.Left != nil || cur.Right != nil)) {
			return false
		}

		if cur.Left != nil {
			queue[r] = cur.Left
			r++
		}
		if cur.Right != nil {
			queue[r] = cur.Right
			r++
		}

		if cur.Left == nil || cur.Right == nil {
			leaf = true
		}
	}

	return true
}
