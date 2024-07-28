package e6247998294f2c933906fdedbc6e6a

import "strings"

// https://www.nowcoder.com/practice/92e6247998294f2c933906fdedbc6e6a

/* 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
*
*
* @param s string字符串
* @return string字符串一维数组
 */
func generatePermutation1(s string) []string {
	set := make(map[string]struct{})
	f1([]rune(s), 0, &strings.Builder{}, &set)
	res := make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	return res
}

// s[i...] 之前决定的路径path，set收集结果时去重
func f1(s []rune, i int, path *strings.Builder, set *map[string]struct{}) {
	if i == len(s) {
		(*set)[path.String()] = struct{}{}
	} else {
		path.WriteRune(s[i])
		f1(s, i+1, path, set)
		// 记录path的内容
		tmpPath := path.String()
		// 重置路径，删除最后一个字符
		path.Reset()
		path.WriteString(tmpPath[:len(tmpPath)-1])
		f1(s, i+1, path, set)
	}
}

func generatePermutation(s string) []string {
	set := make(map[string]struct{})
	f2([]rune(s), 0, &[]rune{}, 0, &set)
	res := make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	return res
}

func f2(s []rune, i int, path *[]rune, size int, set *map[string]struct{}) {
	if i == len(s) {
		(*set)[string((*path)[:size])] = struct{}{}
	} else {
		(*path)[size] = s[i]
		f2(s, i+1, path, size+1, set)
		f2(s, i+1, path, size, set)
	}
}
