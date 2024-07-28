package leetcode105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil || len(preorder) != len(inorder) {
		return nil
	}
	m := make(map[int]int)
	for idx, val := range inorder {
		m[val] = idx
	}
	return f(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, m)
}

func f(pre []int, l1, r1 int, in []int, l2, r2 int, m map[int]int) *TreeNode {
	if l1 > r1 {
		return nil
	}
	head := &TreeNode{Val: pre[l1]}
	if l1 == r1 {
		return head
	}

	k, _ := m[pre[l1]]
	head.Left = f(pre, l1+1, l1+k-l2, in, l2, k-1, m)
	head.Right = f(pre, l1+k-l2+1, r1, in, k+1, r2, m)

	return head
}
