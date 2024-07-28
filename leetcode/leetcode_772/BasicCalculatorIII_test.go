package leetcode772

import (
	"fmt"
	"sync"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []string{"1+1", "2-1+2", "(1+(4+5+2)-3)+(6+8)", "2+(1+14/(4+5*2)-3)+(6+8/4)"}
	results := []int{2, 3, 23, 9}
	for i, test := range tests {
		res := calculate(test)
		if res != results[i] {
			fmt.Printf("Test %d failed, expected %d, got %d\n", i, results[i], res)
		}
	}
}

func TestABC(t *testing.T) {
	ans := 0
	var w sync.WaitGroup

	ch := make(chan int, 1000)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			ch <- 1
		}()
	}
	w.Wait()

	close(ch)

	for c := range ch {
		ans += c
	}
	fmt.Println(ans)
}
