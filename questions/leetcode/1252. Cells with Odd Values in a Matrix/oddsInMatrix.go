package leetcode

func oddCells(n int, m int, indices [][]int) (count int) {
	arr, x, y, i, j, k := make([][]int, n), 0, 0, 0, 0, 0
	for i = 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j = 0; j < m; j++ {
			arr[i][j] = 0
		}
	}
	for i = 0; i < len(indices); i++ {
		x, y = indices[i][0], indices[i][1]
		for j = 0; j < m; j++ {
			arr[x][j]++
		}
		for k = 0; k < n; k++ {
			arr[k][y]++
		}
	}
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			if arr[i][j]%2 != 0 {
				count++
			}
		}
	}
	return count
}
