package leetcode

import "testing"

func TestMaxProduct(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		nums := []int{3, 4, 5, 2}
		want := 12
		got := maxProduct(nums)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
		}
	})

	t.Run("test two", func(t *testing.T) {
		nums := []int{1, 5, 4, 5}
		want := 16
		got := maxProduct(nums)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
		}
	})
}

func TestMaxProductFaster(t *testing.T) {

	t.Run("test one", func(t *testing.T) {
		nums := []int{3, 4, 5, 2}
		want := 12
		got := maxProductFaster(nums)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
		}
	})

	t.Run("test two", func(t *testing.T) {
		nums := []int{1, 5, 4, 5}
		want := 16
		got := maxProductFaster(nums)
		if got != want {
			t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
		}
	})
}
