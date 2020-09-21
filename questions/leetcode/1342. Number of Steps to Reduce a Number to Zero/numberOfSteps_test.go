package leetcode

import "testing"

func TestNumberOfSteps(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := 4
		got := numberOfSteps(8)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := 6
		got := numberOfSteps(14)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test one", func(t *testing.T) {
		want := 12
		got := numberOfSteps(123)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
