package leetcode206

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Ljazz/algorithm/utils"
)

func TestLinkList(t *testing.T) {
	N := 100  // 随机数组最大长度
	V := 1000 // 随机数组每个值，在1~V之间随机
	testTime := 50000
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		// 随机得到一个长度，长度在[0~N-1]
		n := int(rand.Float64() * float64(N))
		// 得到随机数组
		arr := utils.RandomArray(n, V)

		// 构造双向链表
		head := sliceToLinkList(arr)

		newHead := reverseList(head)
		newHead1 := linkedListInversion(head)
		fmt.Println(newHead1)

		var newSlice []int
		for newHead != nil {
			newSlice = append(newSlice, newHead.val)
			newHead = newHead.next
		}

		reverseSlice := utils.ReverseSlice(arr)

		if !utils.SameArray(newSlice, reverseSlice) {
			fmt.Printf("第 %d 次测试出错了！！！\n", i)
		}
	}
	fmt.Println("测试结束")
}

func TestDoubleList(t *testing.T) {
	N := 100  // 随机数组最大长度
	V := 1000 // 随机数组每个值，在1~V之间随机
	testTime := 50000
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		// 随机得到一个长度，长度在[0~N-1]
		n := int(rand.Float64() * float64(N))
		// 得到随机数组
		arr := utils.RandomArray(n, V)

		// 构造双向链表
		head := sliceToDoubleList(arr)

		newHead := reverseDoubleList(head)

		var newSlice []int
		for newHead != nil {
			newSlice = append(newSlice, newHead.val)
			newHead = newHead.next
		}

		reverseSlice := utils.ReverseSlice(arr)

		if !utils.SameArray(newSlice, reverseSlice) {
			fmt.Printf("第 %d 次测试出错了！！！\n", i)
		}
	}
	fmt.Println("测试结束")
}

func sliceToLinkList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var listNode []*ListNode

	for _, v := range arr {
		listNode = append(listNode, NewListNode(v))
	}

	// 处理下标为 0～n 的元素
	for i := 0; i < len(listNode)-1; i++ {
		listNode[i].next = listNode[i+1]
	}
	return listNode[0]
}

func sliceToDoubleList(arr []int) *DoubleListNode {
	if len(arr) == 0 {
		return nil
	}
	var listNode []*DoubleListNode

	for _, v := range arr {
		listNode = append(listNode, NewDoubleListNode(v))
	}

	n := len(arr)
	// 单独处理下标为0与下标为n-1的元素
	// 下标为0的元素的next是下标为1的元素
	// 下标为n-1的元素的last是下标为n-2的元素
	listNode[0].next = listNode[1]
	listNode[n-1].last = listNode[n-2]

	// 处理下标 1～n-2 的元素
	for i := 1; i < n-1; i++ {
		listNode[i].next = listNode[i+1]
	}
	return listNode[0]
}
