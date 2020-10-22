package leetcode

import (
	"testing"
)

var testCases = []struct {
	name     string
	input    []int
	expected int
}{
	{"peak of [0, 1, 0] at index 1", []int{0, 1, 0}, 1},
	{"peak of [0, 2, 1, 0] at index 1", []int{0, 2, 1, 0}, 1},
	{"peak of [0, 10, 5, 2] at index 1", []int{0, 10, 5, 2}, 1},
	{"peak of [3, 4, 5, 1] at index 1", []int{3, 4, 5, 1}, 2},
	{"peak of [24, 69, 100, 99, 79, 78, 67, 36, 26, 19] at index 1", []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, 2},
}

func TestPeakIndexInMountainArray(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := peakIndexInMountainArray(tc.input)
			if actual != tc.expected {
				t.Errorf("for input: %v, expected peak to be at index: %d, but got: %d", tc.input, tc.expected, actual)
			}
		})
	}
}
