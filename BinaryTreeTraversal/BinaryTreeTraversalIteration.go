package binarytreetraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderOneStack(root *TreeNode) []int {
	var res []int
	if root != nil {
		var stack []*TreeNode
		stack = append(stack, root)
		for len(stack) != 0 {
			topIdx := len(stack) - 1
			head := stack[topIdx]       // 栈顶元素
			res = append(res, head.Val) // 将栈顶数据加到结果集中
			stack = stack[:topIdx]      // 删除栈顶元素

			if head.Right != nil {
				stack = append(stack, head.Right)
			}
			if head.Left != nil {
				stack = append(stack, head.Left)
			}
		}
	}
	return res
}

func preOrderOneStack2(root *TreeNode) []int {
	var res []int
	if root != nil {
		var stack []*TreeNode
		stack = append(stack, root)
		for len(stack) != 0 {
			topIdx := len(stack) - 1 // 栈顶元素索引
			// 模拟栈弹出栈顶元素
			head := stack[topIdx]  // 栈顶元素
			stack = stack[:topIdx] // 删除栈顶元素

			res = append(res, head.Val)
			if head.Right != nil {
				stack = append(stack, head.Right)
			}
			if head.Left != nil {
				stack = append(stack, head.Left)
			}
		}
	}
	return res
}

func inorderOneStack(root *TreeNode) []int {
	var res []int
	if root != nil {
		var stack []*TreeNode
		for len(stack) != 0 || root != nil {
			if root != nil {
				stack = append(stack, root)
				root = root.Left
			} else {
				topIdx := len(stack) - 1
				root = stack[topIdx]
				stack = stack[:topIdx]

				res = append(res, root.Val)
				root = root.Right
			}
		}
	}
	return res
}

func posorderTwoStacks(head *TreeNode) []int {
	var res []int
	if head != nil {
		var stack []*TreeNode
		var collect []*TreeNode

		stack = append(stack, head)
		for len(stack) != 0 {
			// pop
			topIdx := len(stack) - 1
			head = stack[topIdx]
			stack = stack[:topIdx]

			collect = append(collect, head)

			if head.Left != nil {
				stack = append(stack, head.Left)
			}
			if head.Right != nil {
				stack = append(stack, head.Right)
			}
		}

		for i := len(collect) - 1; i >= 0; i-- {
			res = append(res, collect[i].Val)
		}
	}
	return res
}

func posorderOneStack(head *TreeNode) []int {
	var res []int
	if head != nil {
		var stack []*TreeNode
		stack = append(stack, head)

		for len(stack) != 0 {
			cur := stack[len(stack)-1]

			if cur.Left != nil && head != cur.Left && head != cur.Right {
				stack = append(stack, cur.Left)
			} else if cur.Right != nil && head != cur.Right {
				stack = append(stack, cur.Right)
			} else {
				res = append(res, cur.Val)
				topIdx := len(stack) - 1
				head = stack[topIdx]
				stack = stack[:topIdx]
			}
		}
	}
	return res
}
