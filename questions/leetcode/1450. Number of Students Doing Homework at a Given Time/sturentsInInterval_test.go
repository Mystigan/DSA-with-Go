package leetcode

import "testing"

func TestBusyStudent(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		start := []int{1, 2, 3}
		end := []int{3, 2, 7}
		want := 1
		got := busyStudent(start, end, 4)
		if got != want {
			t.Errorf("got: %v, want : %v, startTime: %v, endTime: %v", got, want, start, end)
		}
	})

	t.Run("test two", func(t *testing.T) {
		start := []int{4}
		end := []int{4}
		want := 1
		got := busyStudent(start, end, 4)
		if got != want {
			t.Errorf("got: %v, want : %v, startTime: %v, endTime: %v", got, want, start, end)
		}
	})

	t.Run("test three", func(t *testing.T) {
		start := []int{4}
		end := []int{4}
		want := 0
		got := busyStudent(start, end, 5)
		if got != want {
			t.Errorf("got: %v, want : %v, startTime: %v, endTime: %v", got, want, start, end)
		}
	})
}
