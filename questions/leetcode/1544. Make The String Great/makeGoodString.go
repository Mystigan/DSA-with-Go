package leetcode

func makeGood(s string) string {
	stack := []byte{}
	for i := 0; i <= len(s)-1; i++ {
		if len(stack) != 0 && (s[i] == stack[len(stack)-1]+32 || s[i] == stack[len(stack)-1]-32) {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
