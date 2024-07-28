package edfe05a1d45c4ea89101d936cac32469

// 小和问题
// 测试链接 : https://www.nowcoder.com/practice/edfe05a1d45c4ea89101d936cac32469

var help1 []int

func smallSum(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	m := l + (r-l)/2
	part1 := smallSum(arr, l, m)
	part2 := smallSum(arr, m+1, r)
	part3 := merge1(arr, l, m, r)
	return part3 + part2 + part1
}

func merge1(arr []int, l, m, r int) int {
	help1 = make([]int, len(arr))

	ans := 0
	// 4 6 7 7  5     6    6    8
	// i ->  m  j(m+1) ->       r
	// sum = 0
	for i, j, sum := l, m+1, 0; j <= r; j++ {
		for i <= m && arr[i] <= arr[j] {
			sum += arr[i]
			i++
		}
		ans += sum
	}

	i := l
	a := l
	b := m + 1
	for a <= m && b <= r {
		if arr[a] <= arr[b] {
			help1[i] = arr[a]
			a++
		} else {
			help1[i] = arr[b]
			b++
		}
		i++
	}
	for a <= m {
		help1[i] = arr[a]
		a++
		i++
	}
	for b <= r {
		help1[i] = arr[b]
		b++
		i++
	}

	for i := l; i <= r; i++ {
		arr[i] = help1[i]
	}

	return ans
}
