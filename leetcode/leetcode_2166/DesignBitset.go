package leetcode2166

import (
	"strconv"
	"strings"
)

// 位图的实现
// Bitset是一种能以紧凑形式存储位的数据结构
// Bitset(int n) : 初始化n个位，所有位都是0
// void fix(int i) : 将下标i的位上的值更新为1
// void unfix(int i) : 将下标i的位上的值更新为0
// void flip() : 翻转所有位的值
// boolean all() : 是否所有位都是1
// boolean one() : 是否至少有一位是1
// int count() : 返回所有位中1的数量
// String toString() : 返回所有位的状态

type Bitset struct {
	set     []int
	size    int
	zeros   int
	ones    int
	reverse bool
}

func Constructor(size int) Bitset {
	return Bitset{
		set:     make([]int, (size+31)/32),
		size:    size,
		zeros:   size,
		ones:    0,
		reverse: false,
	}
}

func (this *Bitset) Fix(idx int) {
	index := idx / 32
	bit := idx % 32

	if !this.reverse {
		// 位图所有位的状态，维持原有含义
		// 0 : 不存在
		// 1 : 存在
		if (this.set[index] & (1 << bit)) == 0 {
			this.zeros--
			this.ones++
			this.set[index] |= (1 << bit)
		}
	} else {
		// 位图所有位的状态，翻转了
		// 0 : 存在
		// 1 : 不存在
		if (this.set[index] & (1 << bit)) != 0 {
			this.zeros--
			this.ones++
			this.set[index] ^= (1 << bit)
		}
	}
}

func (this *Bitset) Unfix(idx int) {
	index := idx / 32
	bit := idx % 32
	if !this.reverse {
		if (this.set[index] & (1 << bit)) != 0 {
			this.ones--
			this.zeros++
			this.set[index] ^= (1 << bit)
		}
	} else {
		if (this.set[index] & (1 << bit)) == 0 {
			this.ones--
			this.zeros++
			this.set[index] |= (1 << bit)
		}
	}
}

func (this *Bitset) Flip() {
	this.reverse = !this.reverse
	tmp := this.zeros
	this.zeros = this.ones
	this.ones = tmp
}

func (this *Bitset) All() bool {
	return this.ones == this.size
}

func (this *Bitset) One() bool {
	return this.ones > 0
}

func (this *Bitset) Count() int {
	return this.ones
}

func (this *Bitset) ToString() string {
	var builder strings.Builder
	for i, k := 0, 0; i < this.size; k++ {
		number := this.set[k]
		for j := 0; j < 32 && i < this.size; j++ {
			status := (number >> j) & 1
			if this.reverse {
				status ^= 1
			} else {
				status ^= 0
			}
			builder.WriteString(strconv.Itoa(int(status)))
			i++
		}
	}
	return builder.String()
}
