package llstack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	st := InitStack()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)

	t.Run("test Push", func(t *testing.T) {
		want := []interface{}{4, 3, 2, 1}
		got := []interface{}{st.Head.Data}
		curr := st.Head
		for curr.Next != nil {
			curr = curr.Next
			got = append(got, curr.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Pop", func(t *testing.T) {
		want := interface{}(4)
		got, _ := st.Pop()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Peek", func(t *testing.T) {
		want := interface{}(3)
		got, _ := st.Peek()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Size", func(t *testing.T) {
		want := 3
		got := st.Size()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test IsEmpty", func(t *testing.T) {
		want := false
		got := st.IsEmpty()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
