package binarytreetraversal

func preOrderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preOrderTraversal(root.Left)...)
	res = append(res, preOrderTraversal(root.Right)...)
	return res
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

func posorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, posorderTraversal(root.Left)...)
	res = append(res, posorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}
