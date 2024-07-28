package towerofhanoi

import "fmt"

func hanoi(n int) {
	if n > 0 {
		f(n, "左", "右", "中")
	}
}

func f(i int, from, to, other string) {
	if i == 1 {
		fmt.Printf("移动圆盘 1 从 %s 到 %s\n", from, to)
	} else {
		f(i-1, from, other, to)
		fmt.Printf("移动圆盘 %d 从 %s 到 %s\n", i, from, to)
		f(i-1, other, to, from)
	}
}
