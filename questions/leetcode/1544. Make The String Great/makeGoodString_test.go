package leetcode

import "testing"

func TestMakeGood(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "leetcode"
		got := makeGood("leEeetcode")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := ""
		got := makeGood("abBAcC")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
