package leetcode

import "testing"

func TestUniqueMorseRepresentations(t *testing.T) {
	want := 2
	got := uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"})
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
