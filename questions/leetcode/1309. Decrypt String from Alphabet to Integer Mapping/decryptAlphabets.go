package leetcode

import (
	"strconv"
	"strings"
)

func freqAlphabets(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if (i == len(s)-1 || i == len(s)-2 || s[i+2] != '#') && s[i] != '#' {
			b, _ := strconv.Atoi(string(s[i]))
			b1 := byte(b)
			sb.WriteByte(96 + b1)
			continue
		}
		b, _ := strconv.Atoi(s[i : i+2])
		b1 := byte(b)
		sb.WriteByte(96 + b1)
		i += 2
	}
	return sb.String()
}
