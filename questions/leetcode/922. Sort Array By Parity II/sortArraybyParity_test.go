package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArrayByParityII(t *testing.T) {
	want := []int{4, 5, 2, 7}
	got := sortArrayByParityII([]int{4, 2, 5, 7})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
