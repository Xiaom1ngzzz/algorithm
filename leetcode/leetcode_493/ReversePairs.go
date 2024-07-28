package leetcode493

// 翻转对

var help []int

func reversePairs(arr []int) int {
	help = make([]int, len(arr))
	return counts(arr, 0, len(arr)-1)
}

func counts(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	m := (l + r) / 2
	return counts(arr, l, m) + counts(arr, m+1, r) + merge3(arr, l, m, r)
}

func merge3(arr []int, l, m, r int) int {
	ans := 0
	for i := l; i <= m; i++ {
		j := m + 1
		sum := 0
		for j <= r && (arr[i] > 2*arr[j]) {
			sum += 1
			j++
		}
		ans += sum
	}

	// 4 6 7 7  5 8 9 10
	// a        b
	i := l
	a := l
	b := m + 1
	for a <= m && b <= r {
		if arr[a] < arr[b] {
			help[i] = arr[a]
			a++
		} else {
			help[i] = arr[b]
			b++
		}
		i++
	}
	for a <= m {
		help[i] = arr[a]
		a++
		i++
	}
	for b <= r {
		help[i] = arr[b]
		b++
		i++
	}
	for i := l; i <= r; i++ {
		arr[i] = help[i]
	}

	return ans
}
