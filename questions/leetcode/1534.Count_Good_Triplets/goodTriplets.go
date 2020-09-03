package leetcode

func countGoodTriplets(arr []int, a int, b int, c int) (count int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if (abs(arr[i]-arr[j]) <= a) && (abs(arr[j]-arr[k]) <= b) && (abs(arr[i]-arr[k]) <= c) {
					count++
				}
			}
		}
	}
	return count
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
