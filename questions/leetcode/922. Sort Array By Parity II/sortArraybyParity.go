package leetcode

func sortArrayByParityII(A []int) []int {
	//// Brute Force solution
	// for i,v := range A {
	//     // if (i%2==0 && v%2==0) || (i%2!=0 && v%2!=0) {
	//     //     continue
	//     // }
	//     if i%2==0 {
	//         if v%2==0{continue}
	//         for j:=i+1; j<len(A); j++ {
	//             if A[j]%2==0 {
	//                 A[i], A[j] = A[j], A[i]
	//                 break
	//             }
	//         }
	//     } else {
	//         if v%2!=0{continue}
	//         for j:=i+1; j<len(A); j++ {
	//             if A[j]%2!=0 {
	//                 A[i], A[j] = A[j], A[i]
	//                 break
	//             }
	//         }
	//     }
	// }
	// return A

	//// two pointer solution
	// res := make([]int, len(A))
	// start,end := 0,1
	// for _,v := range A {
	//     if v%2==0 {
	//         res[start] = v
	//         start += 2
	//     } else {
	//         res[end] = v
	//         end += 2
	//     }
	// }
	// return res

	// In place swap, most efficient
	for i, j := 0, 1; i < len(A); i += 2 {
		if A[i]&1 == 0 {
			continue
		}
		for ; A[j]&1 != 0; j += 2 {
		}
		A[i], A[j] = A[j], A[i]
	}
	return A
}
