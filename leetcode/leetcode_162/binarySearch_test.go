package leetcode162

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/Ljazz/algorithm/utils"
)

func TestFindNumber(t *testing.T) {
	N := 100  // 随机数组最大长度
	V := 1000 // 随机数组每个值，在1~V之间随机
	testTime := 50000
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		// 随机得到一个长度，长度在[0~N-1]
		n := int(rand.Float64() * float64(N))
		// 得到随机数组
		arr := utils.RandomArray(n, V)
		sort.SliceStable(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		num := int(rand.Float64() * float64(V))
		if right(arr, num) != exist(arr, num) {
			fmt.Printf("第 %d 次测试出错了！！！\n", i)
			fmt.Printf("数组：%v，目标值：%d\n", arr, num)
		} /* else {
			fmt.Printf("第 %d 次测试通过了！！！\n", i)
		}*/
	}
	fmt.Println("测试结束")
}

func TestFindLeft(t *testing.T) {
	N := 100  // 随机数组最大长度
	V := 1000 // 随机数组每个值，在1~V之间随机
	testTime := 50000
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		// 随机得到一个长度，长度在[0~N-1]
		n := int(rand.Float64() * float64(N))
		// 得到随机数组
		arr := utils.RandomArray(n, V)
		sort.SliceStable(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		num := int(rand.Float64() * float64(V))
		if rightFindLeft(arr, num) != findLeft(arr, num) {
			fmt.Printf("第 %d 次测试出错了！！！\n", i)
			fmt.Printf("数组：%v，目标值：%d\n", arr, num)
		} /* else {
			fmt.Printf("第 %d 次测试通过了！！！\n", i)
		}*/
	}
	fmt.Println("测试结束")
}

func TestFindRight(t *testing.T) {
	N := 100  // 随机数组最大长度
	V := 1000 // 随机数组每个值，在1~V之间随机
	testTime := 50000
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		// 随机得到一个长度，长度在[0~N-1]
		n := int(rand.Float64() * float64(N))
		// 得到随机数组
		arr := utils.RandomArray(n, V)
		sort.SliceStable(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		num := int(rand.Float64() * float64(V))
		if rightFindRight(arr, num) != findRight(arr, num) {
			fmt.Printf("第 %d 次测试出错了！！！\n", i)
			fmt.Printf("数组：%v，目标值：%d\n", arr, num)
		} /* else {
			fmt.Printf("第 %d 次测试通过了！！！\n", i)
		}*/
	}
	fmt.Println("测试结束")
}

func TestFindPeakElement(t *testing.T) {
	N := 100  // 随机数组最大长度
	V := 1000 // 随机数组每个值，在1~V之间随机
	testTime := 50000
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		// 随机得到一个长度，长度在[0~N-1]
		n := int(rand.Float64()*float64(N)) + 1
		// 得到随机数组
		arr := utils.RandomArray(n, V)
		num := int(rand.Float64() * float64(V))

		flag := false
		allPeakElement := rightFindPeakElement(arr)
		peakElement := findPeakElement(arr)
		for _, elem := range allPeakElement {
			if elem == peakElement {
				flag = true
			}
		}

		if flag {
			fmt.Printf("第 %d 次测试通过了！！！\n", i)
		} else {
			fmt.Printf("第 %d 次测试出错了！！！\n", i)
			fmt.Printf("数组：%v，目标值：%d\n", arr, num)
		}
	}
	fmt.Println("测试结束")
}
