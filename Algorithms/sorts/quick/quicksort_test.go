package quick

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
		"random array",
		[]int{3, 2, 4, 5, 1},
		[]int{1, 2, 3, 4, 5},
	},
	{
		"sorted array",
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
	},
	{
		"reverse sorted array(worst case)",
		[]int{5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5},
	},
	{
		"empty array",
		[]int{},
		[]int{},
	},
}

func TestQuickSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			want := tc.output
			got := Sort(tc.input, 0, len(tc.input)-1)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got: %v, want: %v", got, want)
			}
		})
	}
}
