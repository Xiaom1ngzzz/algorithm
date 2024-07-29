package eatgrass

// 草一共有n的重量，两只牛轮流吃草，A牛先吃，B牛后吃
// 每只牛在自己的回合，吃草的重量必须是4的幂，1、4、16、64....
// 谁在自己的回合正好把草吃完谁赢，根据输入的n，返回谁赢

func win1(n int) string {
	return f(n, "A")
}

// rest ：还剩下多少草
// cur ：当前选手的名字
// 返回 ： 还剩rest份草，当前选手是cur，按照题目说，返回最终谁赢
func f(rest int, cur string) string {
	enemy := "A"
	if cur == "A" {
		enemy = "B"
	}

	if rest < 5 {
		if rest == 0 || rest == 2 {
			return enemy
		}
		return cur
	}

	pick := 1
	for pick <= rest {
		if f(rest-pick, enemy) == cur {
			return cur
		}
		pick *= 4
	}
	return enemy
}

func win2(n int) string {
	if n%5 == 0 || n%5 == 2 {
		return "B"
	}
	return "A"
}
