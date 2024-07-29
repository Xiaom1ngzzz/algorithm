package eatgrass

import (
	"testing"
)

func TestEatGrass(t *testing.T) {
	count := 50
	for i := 0; i < count; i++ {
		res1 := win1(i)
		res2 := win1(i)

		if res1 != res2 {
			t.Error("出错了")
		}
	}
}
