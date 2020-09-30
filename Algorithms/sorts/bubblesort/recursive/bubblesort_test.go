package bubblesort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("test for empty input", func(t *testing.T) {
		want := []int(nil)
		got, _ := BubbleSort([]int{}, "asc")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test ascending order sort", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5}
		got, _ := BubbleSort([]int{5, 4, 3, 2, 1}, "asc")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test descending order sort", func(t *testing.T) {
		want := []int{5, 4, 3, 2, 1}
		got, _ := BubbleSort([]int{1, 2, 3, 4, 5}, "desc")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test invalid ordering flag", func(t *testing.T) {
		want := []int(nil)
		got, err := BubbleSort([]int{1, 2, 3, 4, 5}, "descending")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
