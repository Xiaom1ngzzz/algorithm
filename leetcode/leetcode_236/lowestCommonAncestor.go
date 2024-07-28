package leetcode236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1. 彼此包含
// 2. 分了两树（即，在两个子树上）

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		// 遇到空，或者p，或者q，直接返回
		return root
	}

	// 左树搜索p和q，遇到空，或者p，或者q，直接返回
	l := lowestCommonAncestor(root.Left, p, q)
	// 右树搜索p和q，遇到空，或者p，或者q，直接返回
	r := lowestCommonAncestor(root.Right, p, q)

	if l != nil && r != nil {
		// 左树也搜到，右树也搜到，返回root
		return root
	}
	if l == nil && r == nil {
		// 都没有搜到，返回空
		return nil
	}

	// l和r一个为空，一个不为空，返回不为空的那个
	if l != nil {
		return l
	}
	return r
}
