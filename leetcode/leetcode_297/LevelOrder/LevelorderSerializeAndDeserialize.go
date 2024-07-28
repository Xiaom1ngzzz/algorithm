package levelorder

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

const MAXN = 10001

var queue [MAXN]*TreeNode

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var b strings.Builder

	l, r := 0, 0
	if root != nil {
		b.WriteString(fmt.Sprintf("%d,", root.Val))
		queue[r] = root
		r++
		for l < r {
			node := queue[l]
			l++
			if node.Left != nil {
				b.WriteString(fmt.Sprintf("%d,", node.Left.Val))
				queue[r] = node.Left
				r++
			} else {
				b.WriteString("#,")
			}

			if node.Right != nil {
				b.WriteString(fmt.Sprintf("%d,", node.Right.Val))
				queue[r] = node.Right
				r++
			} else {
				b.WriteString("#,")
			}
		}
	}
	return b.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	nodes := strings.Split(data, ",")
	index := 0
	root := generate(nodes[index])
	index++
	l, r := 0, 0
	queue[r] = root
	r++
	for l < r {
		cur := queue[l]
		l++
		cur.Left = generate(nodes[index])
		index++
		cur.Right = generate(nodes[index])
		index++
		if cur.Left != nil {
			queue[r] = cur.Left
			r++
		}
		if cur.Right != nil {
			queue[r] = cur.Right
			r++
		}
	}
	return root
}

func generate(val string) *TreeNode {
	if val == "#" {
		return nil
	}
	iVal, _ := strconv.Atoi(val)
	return &TreeNode{Val: iVal}
}
