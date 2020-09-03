package leetcode

import "testing"

func TestCountGoodTriplets(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		nums := []int{3, 0, 1, 1, 9, 7}
		want := 4
		got := countGoodTriplets(nums, 7, 2, 3)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
		}
	})

	t.Run("Second Test", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 3}
		want := 0
		got := countGoodTriplets(nums, 0, 0, 1)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
		}
	})
}

func TestAbs(t *testing.T) {
	t.Run("first abs Test", func(t *testing.T) {
		num := 3
		want := 3
		got := abs(num)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, num)
		}
	})

	t.Run("Second abs Test", func(t *testing.T) {

		num := -3
		want := 3
		got := abs(num)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, num)
		}
	})
}
