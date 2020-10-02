package selection

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	t.Run("test ascending", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5}
		got, _ := Sort([]int{5, 3, 4, 2, 1}, "asc")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test descending", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5}
		got, _ := Sort([]int{2, 3, 1, 5, 4}, "Desc")
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test invalid order", func(t *testing.T) {
		want := []int(nil)
		got, err := Sort([]int{2, 3, 1, 5, 4}, "descen")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

}
