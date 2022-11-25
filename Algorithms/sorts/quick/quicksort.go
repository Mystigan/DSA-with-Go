package quick

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Sort takes in a slice of integers, its start and end indices and returns a slice sorted using the quick sort algorithm.
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

	// get a random index in the range [start, end]
	randomPivotIndex := start + rand.Intn(end - start + 1)

	// swap it with end value
	input[randomPivotIndex], input[end] = input[end], input[randomPivotIndex]

	// accumulate lesser values at front
	for i := start; i < end; i++ {
		if input[i] <= input[end] {
			input[partitionIndex], input[i] = input[i], input[partitionIndex]
			partitionIndex++
		}
	}
	input[partitionIndex], input[end] = input[end], input[partitionIndex]
	return partitionIndex
}
