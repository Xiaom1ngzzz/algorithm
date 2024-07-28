package leetcode912

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Ljazz/algorithm/utils"
)

func TestSelectBubbleInsert(t *testing.T) {
	// arr1 := []int{5, 3, 4, 1, 2}
	// selectionSort(arr1)
	// fmt.Println(arr1)

	// arr2 := []int{5, 3, 4, 1, 2}
	// bubbleSort(arr2)
	// fmt.Println(arr2)

	// arr3 := []int{5, 3, 4, 1, 2}
	// insertionSort(arr3)
	// fmt.Println(arr3)

	N := 100  // 随机数组最大长度
	V := 1000 // 随机数组每个值，在1~V之间随机
	testTime := 5000
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		// 随机得到一个长度，长度在[0~N-1]
		n := int(rand.Float64() * float64(N))
		// 得到随机数组

		arr := utils.RandomArray(n, V)
		arr1 := copyArray(arr)
		arr2 := copyArray(arr)
		arr3 := copyArray(arr)
		arr4 := copyArray(arr)
		selectionSort(arr1)
		bubbleSort(arr2)
		insertionSort(arr3)
		res := BubbleSort(arr4)
		if !utils.SameArray(arr1, arr2) || !utils.SameArray(arr1, arr3) {
			fmt.Println("出错了！！！")
		}
		if !utils.SameArray(arr1, res) {
			fmt.Println("出错了")
		} else {
			fmt.Println("排序算法成功")
			fmt.Println(res)
		}
	}
	fmt.Println("测试结束")
}
