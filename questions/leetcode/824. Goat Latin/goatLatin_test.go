package leetcode

import "testing"

func TestToGoatLatin(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
		got := toGoatLatin("I speak Goat Latin")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		want := "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
		got := toGoatLatin("The quick brown fox jumped over the lazy dog")
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
