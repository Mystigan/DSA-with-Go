package inesrtion

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("test ascending", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5}
		got, _ := Sort([]int{3, 4, 2, 1, 5}, "asc")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test descending", func(t *testing.T) {
		want := []int{5, 4, 3, 2, 1}
		got, _ := Sort([]int{1, 4, 2, 3, 5}, "desc")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test invalid sorting order", func(t *testing.T) {
		want := []int(nil)
		got, err := Sort([]int{1, 4, 2, 3, 5}, "descend")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
