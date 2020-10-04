package merge

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5}
		got := Sort([]int{4, 2, 3, 5, 1})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5}
		got := Sort([]int{5, 4, 3, 2, 1})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
