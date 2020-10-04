package leetcode

import "strings"

func stringMatching(words []string) []string {
	result := []string{}
	for _, v := range words {
		count := 0
		for j := 0; j < len(words); j++ {
			if strings.Contains(words[j], v) {
				count++
			}
		}
		if count > 1 {
			result = append(result, v)
		}

	}
	return result
}
