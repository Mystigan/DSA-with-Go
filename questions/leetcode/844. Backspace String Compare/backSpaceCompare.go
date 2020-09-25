package leetcode

// Stack based implementation
func backspaceCompare(S string, T string) bool {
	s1 := []byte{}
	t1 := []byte{}

	for i := 0; i < len(S); i++ {
		if len(s1) != 0 && S[i] == '#' {
			s1 = s1[:len(s1)-1]
			continue
		}
		if S[i] != '#' {
			s1 = append(s1, S[i])
		}
	}

	for i := 0; i < len(T); i++ {
		if len(t1) != 0 && T[i] == '#' {
			t1 = t1[:len(t1)-1]
			continue
		}
		if T[i] != '#' {
			t1 = append(t1, T[i])
		}
	}

	return string(s1) == string(t1)
}

// // T- O(n)  S- O(1) implementation
// func backspaceCompare(S string, T string) bool {
//     for i:=0; i<len(S); i++ {
//         if S[i]=='#' && i>0 {
//             S = S[:i-1] + S[i+1:]
//             i -= 2
//         } else if S[i]=='#' && i==0 {
//             S = S[1:]
//             i = -1
//         }
//     }
//     for i:=0; i<len(T); i++ {
//         if T[i]=='#' && i>0 {
//             T = T[:i-1] + T[i+1:]
//             i -= 2
//         } else if T[i]=='#' && i==0 {
//             T = T[1:]
//             i = -1
//         }
//     }
//     return S==T
// }
