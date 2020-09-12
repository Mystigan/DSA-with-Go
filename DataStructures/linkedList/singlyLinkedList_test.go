package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedLIst(t *testing.T) {
	t.Run("test complete list and functions", func(t *testing.T) {
		list := InitList()
		list.AddAtHead(7)
		list.AddAtHead(2)
		list.AddAtHead(1)
		list.AddAtIndex(3, 0)
		list.RemoveFromIndex(2)
		list.AddAtHead(6)
		list.AddAtTail(4)
		list.AddAtHead(4)
		list.AddAtIndex(5, 0)
		list.AddAtHead(6)

		list.RemoveFromHead()
		list.RemoveItem(4)
		list.Print()

		// want := []interface{}{6, 4, 6, 1, 2, 0, 0, 4}

		want := []interface{}{6, 1, 2, 0, 0, 4}
		got := []interface{}{}
		current := list.Head
		got = append(got, current.Data)
		for current.Next != nil {
			current = current.Next
			got = append(got, current.Data)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test Get function", func(t *testing.T) {
		list := InitList()
		list.AddAtHead(7)
		list.AddAtHead(2)
		list.AddAtHead(1)
		want := interface{}(2)
		got, _ := list.Get(1)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})

	t.Run("test Find function", func(t *testing.T) {
		list := InitList()
		list.AddAtHead(7)
		list.AddAtHead(2)
		list.AddAtHead(1)
		want := interface{}(2)
		got, _ := list.Find(7)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}

	})
}
