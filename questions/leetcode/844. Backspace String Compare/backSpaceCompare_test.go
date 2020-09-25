package leetcode

import "testing"

func TestBackspaceCompare(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := true
		got := backspaceCompare("ab#c", "ad#c")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := true
		got := backspaceCompare("ab##", "c#d#")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test three", func(t *testing.T) {
		want := false
		got := backspaceCompare("a#c", "b")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
