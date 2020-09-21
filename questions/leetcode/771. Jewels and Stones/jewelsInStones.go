package leetcode

import "strings"

func numJewelsInStones(J string, S string) (count int) {
	for i := 0; i < len(J); i++ {
		// S = strings.Replace(S, string(J[i]), "0", -1)
		count += strings.Count(S, string(J[i]))
	}
	return count
}
