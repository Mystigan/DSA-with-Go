package leetcode

func searchRotatedArray(nums []int, target int) int {
	mini, minVal, left, right := 0, nums[0], 0, 0
	for i, v := range nums {
		if v < minVal {
			minVal = v
			mini = i
		}
	}
	n := len(nums)
	if target <= nums[n-1] {
		if target == nums[n-1] {
			return n - 1
		}
		left = mini
		right = n - 1
	} else {
		right = mini - 1
	}
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
