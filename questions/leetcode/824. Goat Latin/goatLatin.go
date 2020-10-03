package leetcode

import "strings"

func toGoatLatin(S string) string {
	str := strings.Fields(S)
	for i, V := range str {
		v := []byte(V)
		if v[0] == 'a' || v[0] == 'A' || v[0] == 'e' || v[0] == 'E' || v[0] == 'i' || v[0] == 'I' || v[0] == 'o' || v[0] == 'O' || v[0] == 'u' || v[0] == 'U' {
			v = append(v, 'm', 'a')
		} else {
			val := v[0]
			v = v[1:]
			v = append(v, val, 'm', 'a')
		}
		for j := 0; j < i+1; j++ {
			v = append(v, 'a')
		}
		str[i] = string(v)
	}
	return strings.Join(str, " ")
}
