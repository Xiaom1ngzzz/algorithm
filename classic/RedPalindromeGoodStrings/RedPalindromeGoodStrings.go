package redpalindromegoodstrings

// 可以用r、e、d三种字符拼接字符串，如果拼出来的字符串中
// 有且仅有1个长度>=2的回文子串，那么这个字符串定义为"好串"
// 返回长度为n的所有可能的字符串中，好串有多少个
// 结果对1000000007取模， 1 <= n <= 10^9
// 示例：
// n = 1, 输出0
// n = 2, 输出3
// n = 3, 输出18

func num1(n int) int {
	path := make([]rune, n)
	return f(&path, 0)
}

func f(path *[]rune, i int) int {
	if i == len(*path) {
		cnt := 0
		for l := 0; l < len(*path); l++ {
			for r := l + 1; r < len(*path); r++ {
				if is(path, l, r) {
					cnt++
				}
				if cnt > 1 {
					return 0
				}
			}
		}
		if cnt == 1 {
			return 1
		}
		return 0
	}
	ans := 0
	(*path)[i] = 'r'
	ans += f(path, i+1)
	(*path)[i] = 'e'
	ans += f(path, i+1)
	(*path)[i] = 'd'
	ans += f(path, i+1)
	return ans
}

func is(s *[]rune, l, r int) bool {
	for l < r {
		if (*s)[l] != (*s)[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func num2(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 3
	}
	if n == 3 {
		return 18
	}
	return (int)((6 * (n + 1)) % 1000000007)
}
