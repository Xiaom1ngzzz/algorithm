package leetcode102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	levels := make(map[*TreeNode]int)

	queue = append(queue, root)
	levels[root] = 0

	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		level, _ := levels[cur]
		if len(ans) == level {
			levelItems := make([]int, 0)
			ans = append(ans, levelItems)
		}
		ans[level] = append(ans[level], cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			levels[cur.Left] = level + 1
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			levels[cur.Right] = level + 1
		}
	}
	return ans
}

const MAXN = 2001

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var queue [MAXN]*TreeNode
	ans := make([][]int, 0)
	l, r := 0, 0

	queue[r] = root
	r++
	for l < r { // 队列中是否有东西 l == r 的时候表示队列中为空了
		size := r - l
		list := make([]int, 0)

		for i := 0; i < size; i++ {
			cur := queue[l]
			l++
			list = append(list, cur.Val)

			if cur.Left != nil {
				queue[r] = cur.Left
				r++
			}

			if cur.Right != nil {
				queue[r] = cur.Right
				r++
			}
		}
		ans = append(ans, list)
	}

	return ans
}
