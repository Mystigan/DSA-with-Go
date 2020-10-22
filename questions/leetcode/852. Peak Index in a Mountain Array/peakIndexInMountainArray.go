package leetcode

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if arr[mid] < arr[mid+1] {
			left = mid
		} else {
			right = mid
		}
	}
	return -1
}
