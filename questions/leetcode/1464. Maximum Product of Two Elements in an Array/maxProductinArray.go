package leetcode

////short and sweet solution
import "sort"

func maxProduct(nums []int) int {
	sort.Ints(nums)
	return ((nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1))
}

//Faster solution
func maxProductFaster(nums []int) int {
	max, miniMax := 0, 0
	for _, v := range nums {
		if v > max {
			miniMax = max
			max = v
			continue
		}
		if v > miniMax {
			miniMax = v
		}
	}
	return (max - 1) * (miniMax - 1)
}
