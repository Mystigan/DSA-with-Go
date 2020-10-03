package leetcode

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := true
		got := canMakeArithmeticProgression([]int{3, 5, 1})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := false
		got := canMakeArithmeticProgression([]int{1, 2, 4})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
