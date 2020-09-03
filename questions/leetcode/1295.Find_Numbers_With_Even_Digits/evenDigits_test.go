package leetcode

import "testing"

func TestFindNumbers(t *testing.T) {
	nums := []int{12, 345, 2, 6, 7896}
	got := findNumbers(nums)
	want := 2
	if got != want {
		t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
	}
}

func TestIsDigitCountEven(t *testing.T) {
	num := 123
	got := isDigitCountEven(num)
	want := false
	if got != want {
		t.Errorf("got: %v, want: %v, given: %v", got, want, num)
	}
}
