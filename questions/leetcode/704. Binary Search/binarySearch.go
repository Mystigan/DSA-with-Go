package leetcode

func search(nums []int, target int) int {
	result := binarySearch(nums, 0, len(nums)-1, target)
	return result
}

func binarySearch(nums []int, left, right, target int) int {
	if right >= left {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		if target < nums[middle] {
			return binarySearch(nums, left, middle-1, target)
		} else {
			return binarySearch(nums, middle+1, right, target)
		}
	}
	return -1
}
