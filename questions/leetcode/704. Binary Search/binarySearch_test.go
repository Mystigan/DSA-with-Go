package leetcode

import "testing"

var testCases = []struct {
	name     string
	input    []int
	target   int
	expected int
}{
	{"element found toward right end of input slice", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
	{"element found toward left end of input slice", []int{-1, 0, 3, 5, 9, 12}, 0, 1},
	{"element not found in input slice", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
}

func TestBinarySearch(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := search(tc.input, tc.target)
			if actual != tc.expected {
				t.Errorf("For input %v, expected %d to be found at %d, but got %d", tc.input, tc.target, tc.expected, actual)
			}
		})
	}
}
