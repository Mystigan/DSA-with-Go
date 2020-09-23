package leetcode

func buildArray(target []int, n int) []string {
	var result []string
	i, index := 1, 0
	for index < len(target) {
		result = append(result, "Push")
		if target[index] == i {
			index++
		} else {
			result = append(result, "Pop")
		}
		i++
	}
	return result
}
