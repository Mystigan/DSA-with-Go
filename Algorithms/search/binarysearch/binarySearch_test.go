package binarysearch

import (
	"testing"
)

var testCases = []struct {
	name   string
	input  []int
	target int
	output int
}{
	{
		"element found at left hand side",
		[]int{1, 2, 3, 4, 5},
		2,
		1,
	},
	{
		"element found at right hand side",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		7,
		6,
	},
	{
		"element not found",
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		10,
		-1,
	},
	{
		"element found at the middle",
		[]int{1, 2, 3, 4, 5},
		3,
		2,
	},
	{
		"searching in an empty slice",
		[]int{},
		3,
		-1,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want := tc.output
			got := BinarySearch(tc.input, tc.target)
			if got != want {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}
}
