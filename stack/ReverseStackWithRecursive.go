package stack

// 用递归函数逆序栈

func reverse(s *Stack) {
	if s.IsEmpty() {
		return
	}
	num := bottomOut(s)
	reverse(s)
	s.Push(num)
}

// 栈底元素移除掉，上面的元素盖下来
// 返回移除的栈底元素
func bottomOut(s *Stack) int {
	ans, _ := s.Pop()
	if s.IsEmpty() {
		return ans.(int)
	}
	last := bottomOut(s)
	s.Push(ans)
	return last
}
