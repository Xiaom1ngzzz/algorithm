package appleminbags

import "math"

// 有装下8个苹果的袋子、装下6个苹果的袋子，一定要保证买苹果时所有使用的袋子都装满
// 对于无法装满所有袋子的方案不予考虑，给定n个苹果，返回至少要多少个袋子
// 如果不存在每个袋子都装满的方案返回-1

func bags1(apple int) int {
	ans := f(apple)
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

// 当前还有rest个苹果，使用的每个袋子必须装满，返回至少几个袋子
func f(rest int) int {
	if rest < 0 {
		return math.MaxInt
	}
	if rest == 0 {
		return 0
	}
	// 使用8规格的袋子，剩余的苹果还需要几个袋子，有可能返回无效解
	p1 := f(rest - 8)
	// 使用6规格的袋子，剩余的苹果还需要几个袋子，有可能返回无效解
	p2 := f(rest - 6)
	if p1 != math.MaxInt {
		p1++
	}
	if p2 != math.MaxInt {
		p2++
	}
	return min(p1, p2)
}

func bags2(apple int) int {
	if (apple & 1) != 0 {
		return -1
	}
	if apple < 18 {
		if apple == 0 {
			return 0
		}

		if apple == 6 || apple == 8 {
			return 1
		}

		if apple == 12 || apple == 14 || apple == 16 {
			return 2
		}
		return -1
	}
	return (apple-18)/8 + 3
}
