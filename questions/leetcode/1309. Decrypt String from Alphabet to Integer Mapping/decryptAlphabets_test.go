package leetcode

import "testing"

func TestFreqAlphabets(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "jkab"
		got := freqAlphabets("10#11#12")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := "abcdefghijklmnopqrstuvwxyz"
		got := freqAlphabets("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
