type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}
func (s stack) Pop() (stack, bool) {
	if len(s) >= 1 {
		return s[:len(s)-1], true
	} else {
		return s, false
	}
}
func (s stack) Peek() (int, bool) {
	if len(s) >= 1 {
		return s[len(s)-1], true
	} else {
		return -1, false
	}
}

func dailyTemperatures(temperatures []int) []int {
	s := make(stack, 0)
	output := make([]int, len(temperatures))

	for i, temp := range temperatures {
		top, ok := s.Peek()
		for ok && temp > temperatures[top] {
			output[top] = i - top
			s, _ = s.Pop()
			top, ok = s.Peek()
		}
		s = s.Push(i)
	}

	return output
}
