package leetcode

import "testing"

func TestMinOperations(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := 2
		got := minOperations([]string{"d1/", "d2/", "../", "d21/", "./"})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := 3
		got := minOperations([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
