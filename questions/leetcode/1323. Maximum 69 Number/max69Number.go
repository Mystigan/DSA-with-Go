package leetcode

import "strconv"

func maximum69Number(num int) int {
	numStr := strconv.Itoa(num)
	var res []byte
	flag := false
	for i := range numStr {
		if numStr[i] == '6' && flag == false {
			res = append(res, '9')
			flag = true
			continue
		}
		res = append(res, numStr[i])
	}
	num, _ = strconv.Atoi(string(res))
	return num
}

//Without libraries
//     pos, add, curr := 1, 0, num
//     for curr > 0 {
//         d := curr%10
//         curr /= 10

//         if d == 6 {
//             add = 3*pos
//         }

//         pos *= 10
//     }
//     return num + add
// }
