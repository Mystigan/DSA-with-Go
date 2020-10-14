package leetcode

import "testing"

var testCases = []struct {
	name     string
	input    int
	expected int
}{
	{"4^0.5", 4, 2},
	{"8^0.5", 8, 2},
	{"9^0.5", 9, 3},
	{"30^0.5", 30, 5},
	{"10000^0.5", 10000, 100},
}

func TestMySqrt(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := mySqrt(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected sqrt of %d to be %d, but got %d", tc.input, tc.expected, actual)
			}
		})
	}
}
