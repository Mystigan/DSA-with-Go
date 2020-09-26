package leetcode

import "testing"

func TestDestCity(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		want := "Sao Paulo"
		got := destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test one", func(t *testing.T) {
		want := "A"
		got := destCity([][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}})
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
