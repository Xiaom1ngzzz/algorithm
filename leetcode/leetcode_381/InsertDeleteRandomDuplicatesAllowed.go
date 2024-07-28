package leetcode381

import "math/rand"

type RandomizedCollection struct {
	// 用于存储加进来的 value与其在动态数组中的index
	// key ==> value
	// value ==> index 集合
	m map[int][]int
	// 数据
	arr []int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		m:   make(map[int][]int),
		arr: make([]int, 0),
	}
}

func (this *RandomizedCollection) Insert(val int) bool {
	// 将数据添加到动态数组中
	this.arr = append(this.arr, val)
	// 判断 val 之前有没有加入到动态数组中过
	vals, ok := this.m[val]
	if ok {
		// val 值已经存在过，将其索引添加到对应的索引集合中
		vals = append(vals, len(this.arr)-1)
		this.m[val] = vals
		// fmt.Printf("INSERT: %v %v\n", this.arr, this.m)
		return false
	}
	// map中不存在当前val，创建一个索引集合，将当前val的索引加入到索引集合中
	idxs := make([]int, 0)
	idxs = append(idxs, len(this.arr)-1)
	this.m[val] = idxs
	// fmt.Printf("INSERT: %v %v\n", this.arr, this.m)
	return true
}

func (this *RandomizedCollection) Remove(val int) bool {
	// 从 map 中获取当前的val
	vals, ok := this.m[val]
	if !ok {
		// 不存在，直接返回
		return false
	}
	// 动态数组中的最后一个数
	endValue := this.arr[len(this.arr)-1]
	// 当前数与动态数组中最后一个数是否相等
	if val == endValue {
		// 当前数就是动态数组的最后一个
		// 删除 map 中对应的index
		index := 0
		for idx, val := range vals {
			if val == len(this.arr)-1 {
				index = idx
				break
			}
		}
		vals = append(vals[:index], vals[index+1:]...)
	} else {
		// 取索引集合中的第1个
		valAnyIndex := vals[0]
		// 最后一个数的索引集合
		endValueSet := this.m[endValue]
		endValueSet = append(endValueSet, valAnyIndex)
		// 用动态数组的最后一个值覆盖掉当前的值
		this.arr[valAnyIndex] = this.arr[len(this.arr)-1]

		// 删除 endValueSet中的 len(this.arr)-1 值
		index := 0
		for idx, val := range endValueSet {
			if val == len(this.arr)-1 {
				index = idx
				break
			}
		}
		endValueSet = append(endValueSet[:index], endValueSet[index+1:]...)
		this.m[endValue] = endValueSet

		// 删除索引集合中的第1个值
		vals = vals[1:]
	}
	// 删除动态数组的最后一个元素
	this.arr = this.arr[:len(this.arr)-1]

	// fmt.Println(this.arr, this.m)
	// 如果当前key的 value 集合为空，删除
	if len(vals) == 0 {
		delete(this.m, val)
	} else {
		this.m[val] = vals
	}
	return true
}

func (this *RandomizedCollection) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
