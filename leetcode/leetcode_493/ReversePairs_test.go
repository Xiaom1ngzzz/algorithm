package leetcode493

import (
	"fmt"
	"testing"
)

func TestReversePairs(t *testing.T) {
	// arr1 := []int{1, 3, 2, 3, 1}
	// count1 := reversePairs(arr1)
	// if count1 != 2 {
	// 	fmt.Println(count1)
	// 	t.Fatal("出错了")
	// }

	fmt.Println("***********************")
	arr2 := []int{2, 4, 3, 5, 1}
	count2 := reversePairs(arr2)
	if count2 != 3 {
		fmt.Println(count2)
		t.Fatal("出错了")
	}
}
