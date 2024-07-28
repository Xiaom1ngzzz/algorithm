package classic

import (
	"fmt"
	"math"
	"testing"
)

func TestPrintBinary(t *testing.T) {
	num := 7
	printBinary(num)
	// 非负数
	a := 78
	fmt.Println(a)
	printBinary(a)
	fmt.Println("===a===")
	// 负数
	b := -6
	fmt.Println(b)
	printBinary(b)
	fmt.Println("===b===")
	// 直接写二进制的形式定义变量
	c := 0b1001110
	fmt.Println(c)
	printBinary(c)
	fmt.Println("===c===")
	// 直接写十六进制的形式定义变量
	// 0100 -> 4
	// 1110 -> e
	// 0x4e -> 01001110
	d := 0x4e
	fmt.Println(d)
	printBinary(d)
	fmt.Println("===d===")
	// ~、相反数
	fmt.Println(a)
	printBinary(a)
	printBinary(^a)
	e := ^a + 1
	fmt.Println(e)
	printBinary(e)
	fmt.Println("===e===")
	// int、long的最小值，取相反数、绝对值，都是自己
	f := math.MinInt // Integer.MIN_VALUE;
	fmt.Println(f)
	printBinary(f)
	fmt.Println(-f)
	printBinary(-f)
	fmt.Println(^f + 1)
	printBinary(^f + 1)
	fmt.Println("===f===")
	// | & ^
	g := 0b0001010
	h := 0b0001100
	printBinary(g | h)
	printBinary(g & h)
	printBinary(g ^ h)
	fmt.Println("===g、h===")
	// 可以这么写 : int num = 3231 | 6434;
	// 可以这么写 : int num = 3231 & 6434;
	// 不能这么写 : int num = 3231 || 6434;
	// 不能这么写 : int num = 3231 && 6434;
	// 因为 ||、&& 是 逻辑或、逻辑与，只能连接boolean类型
	// 不仅如此，|、& 连接的两侧一定都会计算
	// 而 ||、&& 有穿透性的特点
	fmt.Println("test1测试开始")
	test1 := returnTrue() || returnFalse()
	fmt.Println("test1结果，", test1)
	fmt.Println("test2测试开始")
	test2 := returnTrue() || returnFalse()
	fmt.Println("test2结果，", test2)
	fmt.Println("test3测试开始")
	test3 := returnFalse() && returnTrue()
	fmt.Println("test3结果，", test3)
	fmt.Println("test4测试开始")
	test4 := returnFalse() && returnTrue()
	fmt.Println("test4结果，", test4)
	fmt.Println("===|、&、||、&&===")
	// <<
	i := 0b0011010
	printBinary(i)
	printBinary(i << 1)
	printBinary(i << 2)
	printBinary(i << 3)
	fmt.Println("===i << ===")
	// 非负数 >> >>>，效果一样
	printBinary(i)
	printBinary(i >> 2)
	// printBinary(i >>> 2);
	fmt.Println("===i >> >>>===")
	// 负数 >> >>>，效果不一样
	j := 0b11110000000000000000000000000000
	printBinary(j)
	printBinary(j >> 2)
	// printBinary(j >>> 2);
	fmt.Println("===j >> >>>===")
	// 非负数 << 1，等同于乘以2
	// 非负数 << 2，等同于乘以4
	// 非负数 << 3，等同于乘以8
	// 非负数 << i，等同于乘以2的i次方
	// ...
	// 非负数 >> 1，等同于除以2
	// 非负数 >> 2，等同于除以4
	// 非负数 >> 3，等同于除以8
	// 非负数 >> i，等同于除以2的i次方
	// 只有非负数符合这个特征，负数不要用
	k := 10
	fmt.Println(k)
	fmt.Println(k << 1)
	fmt.Println(k << 2)
	fmt.Println(k << 3)
	fmt.Println(k >> 1)
	fmt.Println(k >> 2)
	fmt.Println(k >> 3)
	fmt.Println("===k===")
}

func returnTrue() bool {
	fmt.Println("进入了returnTrue函数")
	return true
}

func returnFalse() bool {
	fmt.Println("进入了returnFalse函数")
	return false
}
