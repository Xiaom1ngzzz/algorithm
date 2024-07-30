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
