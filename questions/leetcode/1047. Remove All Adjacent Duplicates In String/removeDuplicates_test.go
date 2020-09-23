package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	want := "ca"
	got := removeDuplicates("abbaca")
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
