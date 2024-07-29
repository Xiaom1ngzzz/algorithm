package redpalindromegoodstrings

import (
	"fmt"
	"testing"
)

func TestRedPalindromeGoodStrings(t *testing.T) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("长度为%d, 答案:%d\n", i, num1(i))
	}
	fmt.Println("====================")
	for i := 1; i <= 10; i++ {
		fmt.Printf("长度为%d, 答案:%d\n", i, num2(i))
	}
}
