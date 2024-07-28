package leetcode380

import "math/rand"

// 插入、删除和获取随机元素O(1)时间的结构

type RandomizedSet struct {
	m   map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:   make(map[int]int),
		arr: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok { // 是否存在，存在返回false
		return false
	}
	this.m[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	valIndex, ok := this.m[val]
	if !ok {
		return false
	}

	endValue := this.arr[len(this.arr)-1] // 动态数组中最后一个元素的值
	this.m[endValue] = valIndex
	this.arr[valIndex] = endValue
	delete(this.m, val)
	this.arr = this.arr[:len(this.arr)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
