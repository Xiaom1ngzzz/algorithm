package leetcode297

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

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var b strings.Builder
	serializeTree(root, &b)
	return b.String()
}

func serializeTree(root *TreeNode, builder *strings.Builder) {
	if root == nil {
		builder.WriteString("#,")
	} else {
		builder.WriteString(fmt.Sprintf("%d,", root.Val))
		serializeTree(root.Left, builder)
		serializeTree(root.Right, builder)
	}
}

var cnt int

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")

	cnt = 0

	return deserializeTree(vals)
}

func deserializeTree(vals []string) *TreeNode {
	cur := vals[cnt]
	cnt++

	if cur == "#" {
		return nil
	}
	val, _ := strconv.Atoi(cur)
	head := TreeNode{Val: val}
	head.Left = deserializeTree(vals)
	head.Right = deserializeTree(vals)
	return &head
}
