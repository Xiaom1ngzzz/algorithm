package leetcode190

// 颠倒二进制位

func reverseBits(n uint32) uint32 {
	var ans uint32 = 0
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			ans |= 1 << (31 - i)
		}
	}
	return ans
}

func reverseBits2(n uint32) uint32 {
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	n = (n >> 16) | (n << 16)
	return n
}
