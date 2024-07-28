package leetcode726

import (
	"sort"
	"strconv"
	"strings"
)

var where int

func countOfAtoms(formula string) string {
	where = 0
	m := f([]rune(formula), 0)

	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	ans := strings.Builder{}
	for _, k := range keys {
		num, _ := m[k]
		ans.WriteString(k)
		if num > 1 {
			ans.WriteString(strconv.Itoa(num))
		}
	}
	return ans.String()
}

func f(s []rune, i int) map[string]int {
	// 总表
	ans := make(map[string]int, 0)
	// 之前收集到的名字，历史一部分
	name := strings.Builder{}
	// 之前收集到的有序表，历史一部分
	pre := make(map[string]int, 0)
	// 历史翻几倍
	cnt := 0

	for i < len(s) && s[i] != ')' {
		if s[i] >= 'A' && s[i] <= 'Z' || s[i] == '(' { // 新的开始
			fill(ans, &name, pre, cnt)
			// 归零
			name.Reset()
			pre = nil
			cnt = 0
			if s[i] >= 'A' && s[i] <= 'Z' {
				name.WriteRune(s[i])
				i++
			} else {
				// 遇到 (
				pre = f(s, i+1)
				i = where + 1
			}
		} else if s[i] >= 'a' && s[i] <= 'z' {
			name.WriteRune(s[i])
			i++
		} else {
			cnt = cnt*10 + int(s[i]-'0')
			i++
		}
	}
	fill(ans, &name, pre, cnt)
	where = i
	return ans
}

func fill(ans map[string]int, name *strings.Builder, pre map[string]int, cnt int) {
	if name.Len() > 0 || pre != nil {
		if cnt == 0 {
			cnt = 1
		}
		if name.Len() > 0 {
			key := name.String()
			v, ok := ans[key]
			if ok {
				ans[key] = v + cnt
			} else {
				ans[key] = cnt
			}
		} else {
			for k, v := range pre {
				n, ok := ans[k]
				if ok {
					ans[k] = n + v*cnt
				} else {
					ans[k] = v * cnt
				}
			}
		}
	}
}
