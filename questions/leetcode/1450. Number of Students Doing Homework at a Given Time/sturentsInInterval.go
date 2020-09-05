package leetcode

func busyStudent(startTime []int, endTime []int, queryTime int) (count int) {
	for i := 0; i < len(startTime); i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			count++
		}
	}
	return count
}
