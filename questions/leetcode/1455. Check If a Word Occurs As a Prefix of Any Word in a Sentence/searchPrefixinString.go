package leetcode

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, v := range strings.Split(sentence, " ") {
		if strings.HasPrefix(v, searchWord) {
			return i + 1
		}
		// if len(v)>= len(searchWord) && v[:len(searchWord)]==searchWord{return i+1}
	}
	return -1
}
