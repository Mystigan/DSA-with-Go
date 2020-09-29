package arrstack

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

	t.Run("test Push()", func(t *testing.T) {
		want := Stack{1, 2, 3, 4}
		got := *st

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Pop()", func(t *testing.T) {
		want := interface{}(4)
		got, _ := st.Pop()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Peek()", func(t *testing.T) {
		want := interface{}(3)
		got := st.Peek()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test IsEmpty()", func(t *testing.T) {
		want := false
		got := st.IsEmpty()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Size()", func(t *testing.T) {
		want := 3
		got := st.Size()

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	st.Pop()
	st.Pop()
	st.Pop()
	t.Run("test Pop on empty stack", func(t *testing.T) {
		// want := interface{}(1)
		want := interface{}(nil)
		got, _ := st.Pop()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
