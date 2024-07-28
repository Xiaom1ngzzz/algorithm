package appleminbags

import "math"

func bags1(apple int) int {
	ans := f(apple)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func f(apple int) int {
}
