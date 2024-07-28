package d2707eaf98124f1e8f1d9c18ad487f76

// 不用任何判断语句和比较操作，返回两个数的最大值
// 测试链接 : https://www.nowcoder.com/practice/d2707eaf98124f1e8f1d9c18ad487f76

func flip(n int32) int32 {
	return n ^ 1
}

func sign(n int) int32 {
	return flip(int32(n) >> 31 & 1)
}

func getMaxWithOutJudge(a, b int) int {
	c := a - b
	sa := sign(a) // 0 负数，1 非负数
	sb := sign(b) // 0 负数，1 非负数
	sc := sign(c) // 0 负数，1 非负数

	diffAB := sa ^ sb      // 0 符号一样，1 符号不一样
	sameAB := flip(diffAB) // 0 符号不一样，1 符号一样
	returnA := sa*diffAB + sc*sameAB
	returnB := flip(returnA)
	return a*int(returnA) + b*int(returnB)
}
