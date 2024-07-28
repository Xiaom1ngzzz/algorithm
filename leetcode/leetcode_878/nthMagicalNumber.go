package leetcode878

func nthMagicalNumber(n, a, b int) int {
	abLcm := lcm(a, b)
	ans := 0

	l, r, m := 0, n*min(a, b), 0

	for l <= r {
		m = (l + r) / 2
		if m/a+m/b-m/abLcm >= n {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans % 1000000007
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / gcd(a, b) * b
}

/*
同余原理

(a + b) % m = (a % m + b % m) % m
(a * b) % m = ((a % m) * (b % m)) % m
(a - b) % m = ((a % m - b % m) + m) % m
*/
