package leetcode895

import "math"

type FreqStack struct {
	topTimes   int           // 出现的最大次数
	cntValues  map[int][]int // 每层节点
	valueTimes map[int]int   // 每个数出现次数
}

func Constructor() FreqStack {
	return FreqStack{
		topTimes:   0,
		cntValues:  make(map[int][]int),
		valueTimes: make(map[int]int),
	}
}

func (this *FreqStack) Push(val int) {
	times, ok := this.valueTimes[val]
	curTopTimes := 0
	if !ok {
		this.valueTimes[val] = 1
		curTopTimes = 1
	} else {
		times++
		this.valueTimes[val] = times
		curTopTimes = times
	}

	cntValue, ok := this.cntValues[curTopTimes]
	if !ok {
		cntValue = make([]int, 0)
	}
	cntValue = append(cntValue, val)
	this.cntValues[curTopTimes] = cntValue
	this.topTimes = int(math.Max(float64(this.topTimes), float64(curTopTimes)))
}

func (this *FreqStack) Pop() int {
	topTimeValues, _ := this.cntValues[this.topTimes]
	ans := topTimeValues[len(topTimeValues)-1]
	topTimeValues = topTimeValues[:len(topTimeValues)-1]

	if len(topTimeValues) == 0 {
		delete(this.cntValues, this.topTimes)
		this.topTimes--
	} else {
		this.cntValues[this.topTimes] = topTimeValues
	}

	times, _ := this.valueTimes[ans]
	if times == 1 {
		delete(this.valueTimes, ans)
	} else {
		this.valueTimes[ans] = times - 1
	}

	return ans
}
