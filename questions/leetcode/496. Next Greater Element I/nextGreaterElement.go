package leetcode

//// Brute force solution
// func nextGreaterElement(nums1 []int, nums2 []int) []int {
//     result := []int{}
//     loop:
//         for _,v := range nums1 {
//              j:=0
//             for ;j<len(nums2); j++ {
//                 if nums2[j]==v {
//                     break
//                 }
//             }
//             for ; j<len(nums2); j++ {
//                 if nums2[j] > v{
//                     result = append(result, nums2[j])
//                     continue loop
//                 }
//             }
//             result = append(result, -1)
//         }
//     return result
// }

// Stack based solution
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	data := make(map[int]int, len(nums2))
	var stack []int
	for _, num := range nums2 {
		for len(stack) != 0 && stack[len(stack)-1] < num {
			data[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	ret := make([]int, len(nums1))
	for i, num := range nums1 {
		ret[i] = -1
		if v, ok := data[num]; ok {
			ret[i] = v
		}
	}
	return ret
}
