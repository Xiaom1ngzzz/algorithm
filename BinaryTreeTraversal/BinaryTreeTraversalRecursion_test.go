package binarytreetraversal

import (
	"testing"

	"github.com/Ljazz/algorithm/utils"
)

func TestBinaryTreeTraversalRecursion(t *testing.T) {
	head := TreeNode{Val: 1}
	head.Left = &TreeNode{Val: 2}
	head.Right = &TreeNode{Val: 3}
	head.Left.Left = &TreeNode{Val: 4}
	head.Left.Right = &TreeNode{Val: 5}
	head.Right.Left = &TreeNode{Val: 6}
	head.Right.Right = &TreeNode{Val: 7}

	preRes := preOrderTraversal(&head) // []int{1,2,4,5,3,6,7}
	inRes := inorderTraversal(&head)   // []int{4,2,5,1,6,3,7}
	posRes := posorderTraversal(&head) // []int{4,5,2,6,7,3,1}

	if !utils.SlicesAreEqual(preRes, []int{1, 2, 4, 5, 3, 6, 7}) {
		t.Fatal("先序遍历顺序错误")
	}
	if !utils.SlicesAreEqual(inRes, []int{4, 2, 5, 1, 6, 3, 7}) {
		t.Fatal("中序遍历顺序错误")
	}
	if !utils.SlicesAreEqual(posRes, []int{4, 5, 2, 6, 7, 3, 1}) {
		t.Fatal("后序遍历顺序错误")
	}
}
