package leetcode

import "strings"

func generateTheString(n int) string {
	//// T:100%, mem: 71.1%
	// str := make([]byte, n)
	// for i:=0; i<n-1; i++ {
	//     str[i] = 'a'
	// }
	// if n%2==0 {
	//     str[n-1] = 'b'
	// } else {
	//     str[n-1] = 'a'
	// }
	// return string(str)

	//Alternative solution from discussions.
	res := strings.Repeat("a", n)
	if n&1 == 0 {
		res = "b" + res[1:]
	}
	return res
}
