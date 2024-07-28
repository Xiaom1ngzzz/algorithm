package leetcode231

// Brian Kernighan算法
// 提取出二进制里最右侧的1
// 判断一个整数是不是2的幂

func isPowerOfTwo(n int) bool {
	return n > 0 && n == (n&-n)
}
