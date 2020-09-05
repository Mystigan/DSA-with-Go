package leetcode

import "testing"

func TestCountNegatives(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		nums := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
		got := 8
		want := countNegatives(nums)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
		}
	})

	t.Run("test two", func(t *testing.T) {
		nums := [][]int{{3, 2}, {1, 0}}
		got := 0
		want := countNegatives(nums)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
		}
	})
}
