package leetcode

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)

	t.Run("test Push", func(t *testing.T) {
		want := []int{1, 2}
		got := q.stack
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test Peek", func(t *testing.T) {
		want := 1
		got := q.Peek()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test Pop", func(t *testing.T) {
		want := 1
		got := q.Pop()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test Empty", func(t *testing.T) {
		want := false
		got := q.Empty()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
