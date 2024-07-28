package leetcode222

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return f(root, 1, mostLeft(root, 1))
}

// cur：当前来到的节点
// level：当前来到的节点在第几层
// h：整棵树的高度，不是cur这颗子数的高度
func f(cur *TreeNode, level, h int) int {
	if level == h {
		return 1
	}
	if mostLeft(cur.Right, level+1) == h {
		return 1<<(h-level) + f(cur.Right, level+1, h)
	}
	return 1<<(h-level-1) + f(cur.Left, level+1, h)
}

// 当前节点是cur，并且它在level层
// 返回从cur开始不停往左，能到几层
func mostLeft(cur *TreeNode, level int) int {
	for cur != nil {
		level++
		cur = cur.Left
	}
	return level - 1
}
