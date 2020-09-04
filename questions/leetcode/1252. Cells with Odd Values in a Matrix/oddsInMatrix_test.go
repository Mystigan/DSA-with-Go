package leetcode

import "testing"

func TestOddCells(t *testing.T) {
	want := 6
	got := oddCells(2, 3, [][]int{{0, 1}, {1, 1}})
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
