package leetcode

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name   string
	input  []int
	output []int
}{
	{
		name:   "test one",
		input:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		output: []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
	},
	{
		name:   "test two",
		input:  []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
		output: []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
	},
	{
		name:   "test three",
		input:  []int{10000, 10000},
		output: []int{10000, 10000},
	},
	{
		name:   "test four",
		input:  []int{2, 3, 5, 7, 11, 13, 17, 19},
		output: []int{2, 3, 5, 17, 7, 11, 13, 19},
	},
	{
		name:   "test five",
		input:  []int{10, 100, 1000, 10000},
		output: []int{10, 100, 10000, 1000},
	},
}

func TestSortByBits(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want := tc.output
			got := sortByBits(tc.input)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}
}
