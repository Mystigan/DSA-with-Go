package leetcode

import "strings"

// func sortString(s string) string {
// 	stack := []byte{}
// 	sl := strings.Split(s, "")
// 	sort.Strings(sl)
// 	s = strings.Join(sl, "")

// 	for len(s) > 0 {
// 		stack = append(stack, s[0])
// 		if len(s) == 1 {
// 			break
// 		}
// 		s = s[1:]
// 		for i := 0; i < len(s); i++ {
// 			if stack[len(stack)-1] < s[i] {
// 				stack = append(stack, s[i])
// 				s = s[:i] + s[i+1:]
// 				i--
// 			}
// 		}
//         if len(s)==0 {break}
// 		stack = append(stack, s[len(s)-1])
// 		s = s[:len(s)-1]
// 		for i := len(s) - 1; i >= 0; i-- {
// 			if stack[len(stack)-1] > s[i] {
// 				stack = append(stack, s[i])
// 				s = s[:i] + s[i+1:]
// 			}
// 		}
// 	}
// 	return string(stack)
// }

// Alternate optimised solution from discussions
func sortString(s string) string {
	var sb strings.Builder
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++
	}
	// n:=len(s)
	for sb.Len() < len(s) {
		for i := 0; i < 26; i++ {
			if count[i] > 0 {
				sb.WriteString(string(i + 97))
				count[i]--
				// n--
			}
		}
		for i := 25; i >= 0; i-- {
			if count[i] > 0 {
				sb.WriteString(string(i + 97))
				count[i]--
			}
		}
	}
	return sb.String()
}
