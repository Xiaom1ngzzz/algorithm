package algo

// setAll功能的哈希表
// 测试链接 : https://www.nowcoder.com/practice/7c4559f138e74ceb9ba57d76fd169967

var (
	setAllValue int = 0
	setAllTime  int = -1
	m               = make(map[int][2]int)
	cnt         int = 0
)

func put(k, v int) {
	if val, ok := m[k]; ok {
		val[0] = v
		val[1] = cnt
	} else {
		m[k] = [2]int{v, cnt}
	}
	cnt++
}

func setAll(v int) {
	setAllValue = v
	setAllTime = cnt
	cnt++
}

func get(k int) int {
	val, ok := m[k]

	if !ok {
		return -1
	}

	if val[1] > setAllTime {
		return val[0]
	}

	return setAllValue
}
