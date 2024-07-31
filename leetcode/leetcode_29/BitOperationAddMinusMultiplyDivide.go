package leetcode29

import "math"

// 两数相除
// 不用任何算术运算，只用位运算实现加减乘除
// 代码实现中你找不到任何一个算术运算符

const MIN = math.MinInt

func divide(a int, b int) int {
	if a == MIN && b == MIN {
		// a和b
		return 1
	}
	if a != MIN && b != MIN {
		// a和b都不是整数最小，那么正常去除
		return div(a, b)
	}

	if b == MIN {
		// a不是整数最小，b是整数最小
		return 0
	}
	// a是整数最小，b是-1，返回整数最大，因为题目里明确这么说了
	if b == neg(1) {
		return math.MaxInt
	}
	// a是整数最小，b不是整数最小，b也不是-1
	if b > 0 {
		a = add(a, b)
	} else {
		a = add(a, neg(b))
	}
	ans := div(a, b)
	var offset int
	if b > 0 {
		offset = neg(1)
	} else {
		offset = 1
	}
	return add(ans, offset)
}

// 必须保证a和b都不是整数最小值，返回a除以b的结果
func div(a, b int) int {
	var x int
	if a < 0 {
		x = neg(a)
	} else {
		x = a
	}
	var y int
	if b < 0 {
		y = neg(b)
	} else {
		y = b
	}

	ans := 0
	for i := 30; i >= 0; i = minus(i, 1) {
		if (x >> i) >= y {
			ans |= (1 << i)
			x = minus(x, y<<i)
		}
	}
	if a < 0^b && 0^b < 0 {
		return neg(ans)
	}
	return ans
}

func add(a, b int) int {
	ans := a

	for b != 0 {
		/// ans : a和b无进位相加的结果
		ans := a ^ b
		/// b : a和b相加时的进位信息
		b = (a & b) << 1
		a = ans
	}
	return ans
}

func minus(a, b int) int {
	return add(a, neg(b))
}

func neg(n int) int {
	return add(^n, 1)
}

func multiply(a, b int) int {
	ans := 0

	for b != 0 {
		if (b & 1) != 0 {
			// 考察b当前最右的状态！
			ans = add(ans, a)
		}
		a <<= 1
		b = int(uint(b) >> 1)
	}
	return ans
}
