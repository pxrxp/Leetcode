func reverse(str []byte, start int, end int) {
	left, right := start+1, end-1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
}

func cleanup(str []byte) string {
	result := make([]byte, 0, len(str))
	for _, c := range str {
		if c != '(' && c != ')' {
			result = append(result, c)
		}
	}
	return string(result)
}

func reverseParentheses(s string) string {
	str := []byte(s)
	stack := make([]int, 0, len(s)/2)

	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			reverse(str, stack[len(stack)-1], i)
			stack = stack[:len(stack)-1]
		}
	}

	return cleanup(str)
}
