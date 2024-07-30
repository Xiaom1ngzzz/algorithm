package d88ef50f8dab4850be8cd4b95514bbbd

import (
	"math"
)

// 现在有一个打怪类型的游戏，这个游戏是这样的，你有n个技能
// 每一个技能会有一个伤害，
// 同时若怪物小于等于一定的血量，则该技能可能造成双倍伤害
// 每一个技能最多只能释放一次，已知怪物有m点血量
// 现在想问你最少用几个技能能消灭掉他(血量小于等于0)
// 技能的数量是n，怪物的血量是m
// i号技能的伤害是x[i]，i号技能触发双倍伤害的血量最小值是y[i]
// 1 <= n <= 10
// 1 <= m、x[i]、y[i] <= 10^6
// 测试链接 : https://www.nowcoder.com/practice/d88ef50f8dab4850be8cd4b95514bbbd

const MAXN = 11

var (
	kill  [MAXN]int
	blood [MAXN]int
)

// kill[i]、blood[i]
// n : 一共几个技能
// i : 当前来到了第几号技能
// r : 怪兽目前的剩余血量

func f(n, i, r int) int {
	// 之前的决策已经让怪兽挂了！返回使用了多少个技能
	if r <= 0 {
		return i
	}

	// r > 0
	if i == n {
		// 无效，之前的决策无效
		return math.MaxInt
	}

	ans := math.MaxInt
	for j := i; j < n; j++ {
		swap(i, j)
		if r > blood[i] {
			ans = min(ans, f(n, i+1, r-kill[i]))
		} else {
			ans = min(ans, f(n, i+1, r-(kill[i]*2)))
		}
		swap(i, j)
	}
	return ans
}

func swap(i, j int) {
	tmp := kill[i]
	kill[i] = kill[j]
	kill[j] = tmp

	tmp = blood[i]
	blood[i] = blood[j]
	blood[j] = tmp
}
