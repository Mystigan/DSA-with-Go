package leetcode

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	st := Constructor()
	st.Push(1)
	st.Push(2)

	t.Run("test Push", func(t *testing.T) {
		want := []int{1, 2}
		got := st.queue
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test Top", func(t *testing.T) {
		want := 2
		got := st.Top()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test Pop", func(t *testing.T) {
		want := 2
		got := st.Pop()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test Empty", func(t *testing.T) {
		want := false
		got := st.Empty()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
