package leetcode

import "testing"

func TestGenerateTheString(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "aaa"
		got := generateTheString(3)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := "baaa"
		got := generateTheString(4)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
