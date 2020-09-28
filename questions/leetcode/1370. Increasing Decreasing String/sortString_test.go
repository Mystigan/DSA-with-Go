package leetcode

import "testing"

func TestSortString(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "abccbaabccba"
		got := sortString("aaaabbbbcccc")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := "art"
		got := sortString("rat")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
