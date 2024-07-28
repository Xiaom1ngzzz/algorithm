package leetcode103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MAXN = 2001

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := make([][]int, 0)
	var queue [MAXN]*TreeNode
	l, r := 0, 0

	reverse := false
	queue[r] = root
	r++

	for l < r {
		reverse = !reverse
		size := r - l
		list := make([]int, 0)

		i := 0
		j := 0
		if reverse {
			i = r - 1
			j = -1
		} else {
			i = l
			j = 1
		}
		for k := 0; k < size; k++ {
			cur := queue[i]
			list = append(list, cur.Val)
			i += j
		}

		for i = 0; i < size; i++ {
			cur := queue[l]
			l++

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
