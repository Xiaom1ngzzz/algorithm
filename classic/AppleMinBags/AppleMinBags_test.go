package appleminbags

import (
	"fmt"
	"testing"
)

func TestAppleMinBags(t *testing.T) {
	count := 101
	for i := 1; i < count; i++ {
		res1 := bags1(i)
		res2 := bags2(i)

		fmt.Printf("%d %d\n", res1, res2)

		if res1 != res2 {
			t.Errorf("%d 出错了", i)
		}
	}
}
