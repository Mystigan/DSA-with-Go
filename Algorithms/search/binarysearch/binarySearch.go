package binarysearch

// BinarySearch takes in a slice of integers and a target integer. Searches for the target using binary search algorithm.
// If target is found, returns its index. Otherwise returns -1.
func BinarySearch(nums []int, target int) int {
	// return recursiveSearch(nums, 0, len(nums)-1, target)
	return iterativeSearch(nums, target)
}

func recursiveSearch(nums []int, left, right, target int) int {
	if right >= left {
		middle := left + (right-left)/2
		if target == nums[middle] {
			return middle
		}
		if target < nums[middle] {
			return recursiveSearch(nums, left, middle-1, target)
		}
		return recursiveSearch(nums, middle+1, right, target)
	}
	return -1
}

func iterativeSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		middle := left + (right-left)/2
		if target == nums[middle] {
			return middle
		}
		if target < nums[middle] {
			right = middle - 1
			continue
		}
		left = middle + 1
	}
	return -1
}
