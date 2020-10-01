package leetcode

func reverseString(s []byte) []byte {
	//// two pointer solution
	// lptr, rptr := 0, len(s)-1
	// for rptr > lptr {
	//     s[lptr], s[rptr] = s[rptr], s[lptr]
	//     lptr++
	//     rptr--
	// }

	// efficient solution
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
