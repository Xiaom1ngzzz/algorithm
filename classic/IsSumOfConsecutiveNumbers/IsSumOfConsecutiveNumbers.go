package issumofconsecutivenumbers

// 判断一个数字是否是若干数量(数量>1)的连续正整数的和

func is1(num int) bool {
	sum := 0
	for start := 1; start <= num; start++ {
		sum = start
		for j := start + 1; j <= num; j++ {
			if sum+j > num {
				break
			}
			if sum+j == num {
				return true
			}
			sum += j
		}
	}
	return false
}

func is2(num int) bool {
	return (num & (num - 1)) != 0
}
