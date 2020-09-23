package leetcode

import "testing"

func TestRestoreString(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "leetcode"
		got := restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := "abc"
		got := restoreString("abc", []int{0, 1, 2})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
