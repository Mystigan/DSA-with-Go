package leetcode

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := []int{-1, 3, -1}
		got := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		want := []int{3, -1}
		got := nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
