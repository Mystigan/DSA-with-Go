package leetcode

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	want := []int{1, 3, 6, 10}
	got := runningSum(nums)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
	}
}
