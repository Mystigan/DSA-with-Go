package arrqueue

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := CreateQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	t.Run("test Enqueue()", func(t *testing.T) {
		want := Queue{1, 2, 3, 4}
		got := q
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Peek", func(t *testing.T) {
		want := interface{}(1)
		got := q.Peek()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test IsEmpty", func(t *testing.T) {
		want := false
		got := q.IsEmpty()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Dequeue", func(t *testing.T) {
		want := interface{}(1)
		got, _ := q.Dequeue()
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	t.Run("test Dequeue on empty queue", func(t *testing.T) {
		want := interface{}(nil)
		got, err := q.Dequeue()
		if err != nil {
			fmt.Errorf("error: %v", err)
		}
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
