package leetcode

import "testing"

func TestSubtractProductAndSum(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := 15
		got := subtractProductAndSum(234)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		want := 21
		got := subtractProductAndSum(4421)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
