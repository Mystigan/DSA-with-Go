package leetcode

import "strings"

func numUniqueEmails(emails []string) int {
	strMap := map[string]int{}
	for _, v := range emails {
		index := strings.Index(v, "@")
		strMap[strings.ReplaceAll(strings.Split(v[:index], "+")[0], ".", "")+"@"+v[index+1:]]++
	}
	return len(strMap)
}
