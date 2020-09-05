package leetcode

func busyStudent(startTime []int, endTime []int, queryTime int) (count int) {
	for i := range startTime {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			count++
		}
	}
	return count
}
