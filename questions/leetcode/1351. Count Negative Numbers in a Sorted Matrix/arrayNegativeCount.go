package leetcode

func countNegatives(grid [][]int) (count int) {
	m, n := len(grid), len(grid[0])
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] < 0 {
				count++
			} else {
				break
			}
		}
	}
	return count
}
