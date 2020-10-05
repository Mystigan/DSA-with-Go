package leetcode

import (
	"math/bits"
	"sort"
)

//// Ugly brute force solution
// func sortByBits(arr []int) []int {
//     data := make(map[int][]int, len(arr))
//     for _,v := range arr {
//         if v==0 {
//             data[0] = []int{0}
//             continue
//         }
//         count,n := 0,v
//         for n>0 {
//             count += n & 1
//             n >>= 1
//         }
//         data[count] = append(data[count], v)
//     }
//     temp := []int{}
//     for i := range data {
//         sort.Ints(data[i])
//         temp = append(temp, i)
//     }
//     sort.Ints(temp)
//     result := []int{}
//     for _,v := range temp {
//         result = append(result,data[v]...)
//     }
//     return result
// }

// Solution using sort.Slice
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if x == y {
			return arr[i] < arr[j]
		}
		return x < y
	})
	return arr
}
