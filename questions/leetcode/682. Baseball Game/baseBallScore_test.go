package leetcode

import "testing"

func TestCalPoints(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := 30
		got := calPoints([]string{"5", "2", "C", "D", "+"})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		want := 27
		got := calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
