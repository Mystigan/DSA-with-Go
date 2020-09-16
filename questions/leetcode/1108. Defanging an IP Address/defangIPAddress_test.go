package leetcode

import "testing"

func TestDefangIPaddr(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "1[.]1[.]1[.]1"
		got := defangIPaddr("1.1.1.1")
		if got != want {
			t.Errorf("got: %v, want: %v, given: 1.1.1.1", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		want := "255[.]100[.]50[.]0"
		got := defangIPaddr("255.100.50.0")
		if got != want {
			t.Errorf("got: %v, want: %v, given: 1.1.1.1", got, want)
		}
	})
}
