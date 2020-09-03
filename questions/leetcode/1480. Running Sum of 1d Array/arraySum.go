package leetcode

func runningSum(nums []int) []int {
	var sum int = 0
	var runSum = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		runSum[i] = sum
	}
	return runSum
}
