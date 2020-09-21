package leetcode

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := 3
		got := numJewelsInStones("aA", "aAAbbbb")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		want := 0
		got := numJewelsInStones("z", "ZZZZ")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
