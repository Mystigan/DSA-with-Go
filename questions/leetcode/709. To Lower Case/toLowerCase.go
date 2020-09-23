package leetcode

func toLowerCase(str string) string {
	//// Using in-built function
	//return strings.ToLower(str)

	// Using byte slice
	result := []byte(str)
	for i := range result {
		if result[i] >= 'A' && result[i] <= 'Z' {
			result[i] += 32
		}
	}
	return string(result)
}
