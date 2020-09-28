package leetcode

func minOperations(logs []string) int {
	stack := []string{}
	for i := 0; i < len(logs); i++ {
		if logs[i] == "./" {
			continue
		}
		if logs[i] == "../" {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, logs[i])
	}
	return len(stack)
}
