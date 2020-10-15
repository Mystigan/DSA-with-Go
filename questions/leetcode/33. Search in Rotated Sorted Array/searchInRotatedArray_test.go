package leetcode

import "testing"

var testcases = []struct {
	name   string
	input  []int
	target int
	want   int
}{
	{"0 found in [4,5,6,7,0,1,2]", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	{"3 not found in [4,5,6,7,0,1,2]", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	{"0 not found in [1]", []int{1}, 0, -1},
}

func TestSearchRotatedArray(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := searchRotatedArray(tc.input, tc.target)
			if got != tc.want {
				t.Errorf("Search for %d in %v, expected: %d, got: %d", tc.target, tc.input, tc.want, got)
			}
		})
	}
}
