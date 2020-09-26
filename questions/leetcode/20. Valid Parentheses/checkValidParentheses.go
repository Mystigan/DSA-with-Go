package leetcode

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		if stack[len(stack)-1] == '(' && s[i] == ')' {
			stack = stack[:len(stack)-1]
			continue
		}
		if stack[len(stack)-1] == '{' && s[i] == '}' {
			stack = stack[:len(stack)-1]
			continue
		}
		if stack[len(stack)-1] == '[' && s[i] == ']' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return len(stack) == 0
}
