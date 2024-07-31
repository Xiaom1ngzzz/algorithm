package leetcode906

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	ans := collect()
	for _, v := range ans {
		fmt.Println(v)
	}

	fmt.Println("size: ", len(ans))
}

func TestG(t *testing.T) {
	num := 123

	fmt.Println(even(num))
	fmt.Println(odd(num))
	fmt.Println(evenEnlarge(int64(num)))
	fmt.Println(oddEnlarge(int64(num)))
}

func even(num int) int {
	ans := num

	for num != 0 {
		ans = ans*10 + num%10
		num /= 10
	}
	return ans
}

func odd(num int) int {
	ans := num
	num /= 10

	for num != 0 {
		ans = ans*10 + num%10
		num /= 10
	}
	return ans
}
