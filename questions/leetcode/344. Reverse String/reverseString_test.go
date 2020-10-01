package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := []byte{'o', 'l', 'l', 'e', 'h'}
		got := reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
		got := reverseString([]byte{'H', 'a', 'n', 'n', 'a', 'h'})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
