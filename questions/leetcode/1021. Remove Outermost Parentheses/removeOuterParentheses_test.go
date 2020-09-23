package leetcode

import (
	"testing"
)

func TestRemoveOuterParentheses(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "()()()"
		got := removeOuterParentheses("(()())(())")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := "()()()()(())"
		got := removeOuterParentheses("(()())(())(()(()))")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test three", func(t *testing.T) {
		want := ""
		got := removeOuterParentheses("()()")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
