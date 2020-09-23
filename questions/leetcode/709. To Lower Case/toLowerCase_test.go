package leetcode

import "testing"

func TestToLowerCase(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "hello"
		got := toLowerCase("Hello")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := "lovely"
		got := toLowerCase("LOVEly")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
