package leetcode772

// BasicCalculatorIII https://leetcode.com/problems/basic-calculator-iii/

var where int

func calculate(s string) int {
	where = 0
	return f([]rune(s), 0)
}

// s[i....]开始计算，遇到字符串终止 或者 ) 停止
// 返回 ： 自己负责的这一段，计算的结果
// 返回之前，更新全局的where，为了上游函数知道从哪继续
func f(s []rune, i int) int {
	cur := 0
	numbers := make([]int, 0)
	ops := make([]rune, 0)
	for i < len(s) && s[i] != ')' {
		if s[i] >= '0' && s[i] != '9' {
			cur = cur*10 + int(s[i]-'0')
			i++
		} else if s[i] != '(' {
			// 遇到了运算符 + - * /
			push(&numbers, &ops, cur, s[i])
			i++
			cur = 0
		} else {
			// 遇到了左括号
			cur = f(s, i+1)
			i = where + 1
		}
	}
	push(&numbers, &ops, cur, '+')
	where = i
	return compute(&numbers, &ops)
}

func push(numbers *[]int, ops *[]rune, cur int, op rune) {
	n := len(*numbers)
	if n == 0 || (*ops)[n-1] == '+' || (*ops)[n-1] == '-' {
		*numbers = append(*numbers, cur)
		*ops = append(*ops, op)
	} else {
		topNumber := (*numbers)[n-1]
		topOp := (*ops)[n-1]
		if topOp == '*' {
			*numbers = (*numbers)[:n-1]
			*numbers = append(*numbers, topNumber*cur)
		} else if topOp == '/' {
			*numbers = (*numbers)[:n-1]
			*numbers = append(*numbers, topNumber/cur)
		}
		*ops = (*ops)[:n-1]
		*ops = append(*ops, op)
	}
}

func compute(numbers *[]int, ops *[]rune) int {
	n := len(*numbers)
	ans := (*numbers)[0]
	for i := 1; i < n; i++ {
		if (*ops)[i-1] == '+' {
			ans += (*numbers)[i]
		} else if (*ops)[i-1] == '-' {
			ans -= (*numbers)[i]
		}
	}
	return ans
}
