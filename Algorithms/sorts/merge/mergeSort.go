package merge

func Sort(input []int) []int {
	n := len(input)
	if n < 2 {
		return input
	}
	left := []int{}
	for i := 0; i < n/2; i++ {
		left = append(left, input[i])
	}
	right := []int{}
	for i := n / 2; i < n; i++ {
		right = append(right, input[i])
	}
	Sort(left)
	Sort(right)
	return merge(left, right, input)
}

func merge(left, right, input []int) []int {
	i, j, k, nL, nR := 0, 0, 0, len(left), len(right)
	for i < nL && j < nR {
		if left[i] < right[j] {
			input[k] = left[i]
			i++
		} else {
			input[k] = right[j]
			j++
		}
		k++
	}
	for i < nL {
		input[k] = left[i]
		k++
		i++
	}
	for j < nR {
		input[k] = right[j]
		k++
		j++
	}
	return input
}
