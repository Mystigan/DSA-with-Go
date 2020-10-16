package leetcode

import "testing"

var testCases = []struct {
	name     string
	input    []int
	expected int
}{
	{"min:1", []int{4, 5, 6, 7, 1, 2, 3}, 1},
	{"min:0", []int{5, 6, 7, 0, 1, 2, 3}, 0},
	{"min:3", []int{5, 6, 7, 8, 3, 4}, 3},
	{"min:100", []int{105, 140, 151, 155, 190, 200, 100, 101, 102, 104}, 100},
	{"min:1000", []int{1007, 1010, 1020, 1030, 1031, 1033, 1040, 1050, 1000, 1001, 1002, 1003, 1004}, 1000},
}

func TestFindMin(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findMin(tc.input)
			if actual != tc.expected {
				t.Errorf("For %v, expected min to be: %d, but got: %d", tc.input, tc.expected, actual)
			}
		})
	}
}

func BenchmarkFindMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			findMin(tc.input)
		}
	}
}
