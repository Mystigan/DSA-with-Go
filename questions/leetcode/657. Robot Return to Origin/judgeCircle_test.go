package leetcode

import "testing"

func TestJudgeCircle(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := true
		got := judgeCircle("UD")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := false
		got := judgeCircle("LL")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test three", func(t *testing.T) {
		want := false
		got := judgeCircle("RRDD")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
