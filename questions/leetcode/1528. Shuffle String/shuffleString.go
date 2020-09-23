package leetcode

func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))
	for i, v := range indices {
		result[v] = s[i]
	}
	return string(result)
}
