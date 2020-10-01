package leetcode

import "strings"

// func judgeCircle(moves string) bool {
//     vcount, hcount := 0, 0
//     for _,v := range moves {
//         switch(v) {
//             case 'U':
//                 vcount++
//             case 'D':
//                 vcount--
//             case 'L':
//                 hcount++
//             case 'R':
//                 hcount--
//         }
//     }
//     return vcount==0 && hcount==0
// }

// Alternative solution
func judgeCircle(moves string) bool {
	return strings.Count(moves, "U") == strings.Count(moves, "D") && strings.Count(moves, "L") == strings.Count(moves, "R")
}
