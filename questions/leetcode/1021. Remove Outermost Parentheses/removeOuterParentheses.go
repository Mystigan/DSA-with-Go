package leetcode

func removeOuterParentheses(S string) string {

	sum := 0
	var result []byte
	for a, i := 0, 0; i < len(S); i++ {
		if S[i] == '(' {
			sum++
		} else {
			sum--
			if sum == 0 {
				result = append(result, S[a+1:i]...)
				a = i + 1
			}
		}
	}
	return string(result)
}
