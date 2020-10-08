package quick

// Sort takesin a slice of integers, its start and end indices and returns a slice sorted using the quick sort algorithm.
func Sort(input []int, start, end int) []int {
	if start < end {
		pivotIndex := partition(input, start, end)
		Sort(input, start, pivotIndex-1)
		Sort(input, pivotIndex+1, end)
	}
	return input
}

func partition(input []int, start, end int) int {
	partitionIndex := start
	randomPivotIndex := end
	for i := start; i < end; i++ {
		if input[i] <= input[randomPivotIndex] {
			input[partitionIndex], input[i] = input[i], input[partitionIndex]
			partitionIndex++
		}
	}
	input[partitionIndex], input[randomPivotIndex] = input[randomPivotIndex], input[partitionIndex]
	return partitionIndex
}
