package leetcode

import (
	"reflect"
	"testing"
)

func TestMinStack(t *testing.T) {
	st := Constructor()
	st.Push(-2)
	st.Push(0)
	st.Push(-3)

	t.Run("test Push", func(t *testing.T) {
		want := []int{-2, 0, -3}
		got := st.stack
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Top", func(t *testing.T) {
		want := -3
		got := st.Top()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test GetMin", func(t *testing.T) {
		want := -3
		got := st.GetMin()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Pop", func(t *testing.T) {
		st.Pop()
		want := []int{-2, 0}
		got := st.stack
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test GetMin", func(t *testing.T) {
		want := -2
		got := st.GetMin()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
