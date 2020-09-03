package leetcode

import "testing"

func TestMinTimeToVisitAllPoints(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
		want := 7
		got := minTimeToVisitAllPoints(points)
		if got != want {
			t.Errorf("got: %v, want: %v; given: %v", got, want, points)
		}
	})

	t.Run("Test 1", func(t *testing.T) {
		points := [][]int{{3, 2}, {-2, 2}}
		want := 5
		got := minTimeToVisitAllPoints(points)
		if got != want {
			t.Errorf("got: %v, want: %v; given: %v", got, want, points)
		}
	})
}
