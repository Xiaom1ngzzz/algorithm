package leetcode337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	yes int
	no  int
)

func rob(root *TreeNode) int {
	f(root)
	return max(yes, no)
}

func f(cur *TreeNode) {
	if cur == nil {
		yes = 0
		no = 0
	} else {
		y := cur.Val
		n := 0
		f(cur.Left)
		y += no
		n += max(yes, no)
		f(cur.Right)
		y += no
		n += max(yes, no)
		yes = y
		no = n
	}
}
