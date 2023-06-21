package main

// 有效的括号 stack
func isValid(s string) bool {
	m := make(map[int32]int32)
	m['('] = ')'
	m['['] = ']'
	m['{'] = '}'

	stack := make([]int32, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			continue
		}

		if v == ')' || v == ']' || v == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if m[top] == v {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
