package leetcode

func minTimeToVisitAllPoints(points [][]int) (max int) {
	for i := 0; i < len(points)-1; i++ {
		x1, y1 := points[i][0], points[i][1]
		x2, y2 := points[i+1][0], points[i+1][1]
		x, y := abs(x2-x1), abs(y2-y1)
		if x > y {
			max += x
		} else {
			max += y
		}
	}
	return max
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
