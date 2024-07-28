package leetcode326

// 判断一个整数是不是3的幂

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	if n == 1 {
		return true
	}

	for n > 1 {
		if n%3 == 0 {
			n /= 3
		} else {
			return false
		}
	}
	return true
}

// 如果一个数字是3的某次幂，那么这个数一定只含有3这个质数因子
// 1162261467是int型范围内，最大的3的幂，它是3的19次方
// 这个1162261467只含有3这个质数因子，如果n也是只含有3这个质数因子，那么
// 1162261467 % n == 0
// 反之如果1162261467 % n != 0 说明n一定含有其他因子
func isPowerOfThree2(n int) bool {
	return n > 0 && 1162261467%n == 0
}
