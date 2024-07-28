package leetcode662

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MAXN = 3001

func widthOfBinaryTree(root *TreeNode) int {
	var (
		nq [MAXN]*TreeNode
		iq [MAXN]int
	)

	ans := 1
	l, r := 0, 0

	nq[r] = root
	iq[r] = 1
	r++

	for l < r {
		size := r - l
		ans = max(ans, iq[r-1]-iq[l]+1)

		for i := 0; i < size; i++ {
			node := nq[l]
			id := iq[l]
			l++

			if node.Left != nil {
				nq[r] = node.Left
				iq[r] = id * 2
				r++
			}

			if node.Right != nil {
				nq[r] = node.Right
				iq[r] = id*2 + 1
				r++
			}
		}
	}

	return ans
}
