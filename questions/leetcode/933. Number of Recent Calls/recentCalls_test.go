package leetcode

import "testing"

func TestRecentCounter(t *testing.T) {
	rc := Constructor()
	rc.Ping(1)
	rc.Ping(100)
	t.Run("test one", func(t *testing.T) {
		want := 3
		got := rc.Ping(3001)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := 3
		got := rc.Ping(3002)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
