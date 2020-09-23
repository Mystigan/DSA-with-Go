package leetcode

import "testing"

func TestMaximum69Number(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := 9969
		got := maximum69Number(9669)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := 9999
		got := maximum69Number(9996)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
