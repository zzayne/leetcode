package other

// https://leetcode.cn/problems/valid-parentheses/

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	lst := []rune(s)
	if !isLeft(lst[0]) {
		return false
	}
	if !isRight(lst[len(s)-1]) {
		return false
	}

	match := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{lst[0]}
	for i := 1; i < len(lst); i++ {
		v := lst[i]
		if isLeft(v) {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 || match[stack[len(stack)-1]] != v {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

func isLeft(r rune) bool {
	if r == '(' || r == '{' || r == '[' {
		return true
	}
	return false
}
func isRight(r rune) bool {
	if r == ')' || r == '}' || r == ']' {
		return true
	}
	return false
}
