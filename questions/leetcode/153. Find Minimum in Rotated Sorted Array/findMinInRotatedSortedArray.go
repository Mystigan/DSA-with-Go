package leetcode

func findMin(nums []int) int {
	min := nums[0]
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < min {
			min = nums[mid]
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] < min {
		min = nums[left]
	}
	return min
}
