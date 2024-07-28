package leetcode113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	// 结果集合
	ans := make([][]int, 0)

	if root != nil {
		// 路径
		path := make([]int, 0)
		process(root, targetSum, 0, &path, &ans)
	}

	return ans
}

func process(cur *TreeNode, aim, sum int, path *[]int, ans *[][]int) {
	if cur.Left == nil && cur.Right == nil {
		// 叶节点
		if cur.Val+sum == aim {
			*path = append(*path, cur.Val)
			*ans = append(*ans, append([]int{}, *path...))
			*path = (*path)[:len(*path)-1]
		}
	} else {
		// 非叶节点
		*path = append(*path, cur.Val)
		if cur.Left != nil {
			process(cur.Left, aim, sum+cur.Val, path, ans)
		}
		if cur.Right != nil {
			process(cur.Right, aim, sum+cur.Val, path, ans)
		}
		*path = (*path)[:len(*path)-1]
	}
}
