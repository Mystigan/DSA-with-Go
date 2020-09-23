package leetcode

import (
	"reflect"
	"testing"
)

func TestBuildArray(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		target := []int{1, 3}
		want := []string{"Push", "Push", "Pop", "Push"}
		got := buildArray(target, 3)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		target := []int{1, 2, 3}
		want := []string{"Push", "Push", "Push"}
		got := buildArray(target, 3)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test three", func(t *testing.T) {
		target := []int{2, 3, 4}
		want := []string{"Push", "Pop", "Push", "Push", "Push"}
		got := buildArray(target, 4)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
