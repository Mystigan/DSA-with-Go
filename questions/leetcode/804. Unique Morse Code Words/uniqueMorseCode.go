package leetcode

//// My solution: T: o ms(100%), mem: 2.4mb (41.1%)
// func uniqueMorseRepresentations(words []string) int {
//     letters := map[byte]string{
//         'a': ".-",
//         'b': "-...",
//         'c': "-.-.",
//         'd': "-..",
//         'e': ".",
//         'f': "..-.",
//         'g': "--.",
//         'h': "....",
//         'i': "..",
//         'j': ".---",
//         'k': "-.-",
//         'l': ".-..",
//         'm': "--",
//         'n': "-.",
//         'o': "---",
//         'p': ".--.",
//         'q': "--.-",
//         'r': ".-.",
//         's': "...",
//         't': "-",
//         'u': "..-",
//         'v': "...-",
//         'w': ".--",
//         'x': "-..-",
//         'y': "-.--",
//         'z': "--..",
//     }

//     for i,v := range words {
//         var sb strings.Builder
//         for j:=0; j<len(v); j++ {
//             sb.WriteString(letters[v[j]])
//         }
//         words[i] = sb.String()
//         sb.Reset()
//     }
//     strMap := map[string]int{}
//     for _,v := range words {
//         strMap[v]++
//     }
//     return len(strMap)
// }

// Alternate solution(from discussions) Better memory utilisation: T: 0ms (100%), mem: 2.3mb (81.01%)
func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	ctMap := map[string]int{}
	for _, v := range words {
		str := ""
		for _, val := range v {
			str = str + morse[val-'a']
		}
		ctMap[str]++
	}
	return len(ctMap)
}
