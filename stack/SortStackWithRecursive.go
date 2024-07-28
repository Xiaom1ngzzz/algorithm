package stack

import "math"

// 用递归函数排序栈
// 栈只提供push、pop、isEmpty三个方法
// 请完成无序栈的排序，要求排完序之后，从栈顶到栈底从小到大
// 只能使用栈提供的push、pop、isEmpty三个方法、以及递归函数
// 除此之外不能使用任何的容器，数组也不行
// 就是排序过程中只能用：
// 1) 栈提供的push、pop、isEmpty三个方法
// 2) 递归函数，并且返回值最多为单个整数

func sort(stack *Stack) {
	deep := getDeep(stack)

	for deep > 0 {
		max := getMax(stack, deep)
		k := times(stack, deep, max)
		down(stack, deep, max, k)
		deep -= k
	}
}

// 返回栈的深度
// 不改变栈的数据状况
func getDeep(stack *Stack) int {
	if stack.IsEmpty() {
		return 0
	}
	num, _ := stack.Pop()
	d := getDeep(stack) + 1
	stack.Push(num)
	return d
}

// 从栈当前的顶部开始，往下数deep层
// 返回这deep层里的最大值
func getMax(stack *Stack, deep int) int {
	if deep == 0 {
		return math.MinInt
	}
	num, _ := stack.Pop()
	var restMax int
	restMax = getMax(stack, deep-1)
	m := int(max(float64(restMax), float64(num.(int))))
	stack.Push(num)
	return m
}

// 从栈当前的顶部开始，往下数deep层，已知最大值是max了
// 返回，max出现了几次，不改变栈的数据状况
func times(stack *Stack, deep int, max int) int {
	if deep == 0 {
		return 0
	}
	num, _ := stack.Pop()
	var restTimes int
	restTimes = times(stack, deep-1, max)
	if num.(int) == max {
		restTimes++
	}
	stack.Push(num)
	return restTimes
}

// 从栈当前的顶部开始，往下数deep层，已知最大值是max，出现了k次
// 请把这k个最大值沉底，剩下的数据状况不变
func down(stack *Stack, deep int, max int, k int) {
	if deep == 0 {
		for i := 0; i < k; i++ {
			stack.Push(max)
		}
	} else {
		num, _ := stack.Pop()
		down(stack, deep-1, max, k)
		if num.(int) != max {
			stack.Push(num)
		}
	}
}
