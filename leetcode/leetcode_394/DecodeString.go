package leetcode394

import "strings"

var where int

func decodeString(s string) string {
	where = 0
	return f([]rune(s), 0)
}

// s[i....]开始计算，遇到字符串终止 或者 ) 停止
// 返回 ： 自己负责的这一段，计算的结果
// 返回之前，更新全局的where，为了上游函数知道从哪继续
func f(s []rune, i int) string {
	path := strings.Builder{}

	cnt := 0

	for i < len(s) && s[i] != ']' {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			path.WriteRune(s[i])
			i++
		} else if s[i] >= '0' && s[i] <= '9' {
			cnt = cnt*10 + int(s[i]-'0')
			i++
		} else {
			// 遇到 [
			path.WriteString(get(cnt, f(s, i+1)))
			i = where + 1
			cnt = 0
		}
	}
	where = i
	return path.String()
}

func get(cnt int, str string) string {
	builder := strings.Builder{}
	for i := 0; i < cnt; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
