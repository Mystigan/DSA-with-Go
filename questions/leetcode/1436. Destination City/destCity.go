package leetcode

func destCity(paths [][]string) string {
	strMap := map[string]int{}
	for i := 0; i < len(paths); i++ {
		strMap[paths[i][0]]++
		strMap[paths[i][1]]--
	}
	for i, v := range strMap {
		if v == -1 {
			return i
		}
	}

	// strMap := map[string]string{}
	// for _,v := range paths {
	//     strMap[v[0]] = v[1]
	// }
	// for _,v := range paths {
	//     if _,ok := strMap[v[1]]; !ok {
	//         return v[1]
	//     }
	// }
	return ""
}
