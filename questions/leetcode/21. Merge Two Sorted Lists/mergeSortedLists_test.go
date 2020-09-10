package leetcode

import (
	"reflect"
	"testing"
)

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Length int
	Head   *Node
}

func (ll *List) Add(data int) bool {
	node := &Node{
		Val: data,
	}
	if ll.Head == nil {
		ll.Head = node
		ll.Length++
		return true
	}

	//// Adding at the beginning of the list.
	// node.Next = ll.Head
	// ll.Head = node
	// ll.Length++
	// return true

	// Adding at the end of the list.
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
	ll.Length++
	return true
}

func TestMergeLists(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		l1 := &List{}
		l1.Add(1)
		l1.Add(2)
		l1.Add(4)
		l2 := &List{}
		l2.Add(1)
		l2.Add(3)
		l2.Add(4)

		want := []int{1, 1, 2, 3, 4, 4}

		l3 := mergeLists(l1.Head, l2.Head)
		got := []int{}
		got = append(got, l3.Val)
		for l3.Next != nil {
			l3 = l3.Next
			got = append(got, l3.Val)
		}
		if val := reflect.DeepEqual(got, want); !val {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		l1 := &List{}
		l1.Add(1)
		l1.Add(2)
		l1.Add(4)
		l2 := &List{}
		l2.Add(1)
		l2.Add(3)
		l2.Add(4)
		l2.Add(5)
		l2.Add(6)
		l2.Add(7)

		want := []int{1, 1, 2, 3, 4, 4, 5, 6, 7}

		l3 := mergeLists(l1.Head, l2.Head)
		got := []int{}
		got = append(got, l3.Val)
		for l3.Next != nil {
			l3 = l3.Next
			got = append(got, l3.Val)
		}
		if val := reflect.DeepEqual(got, want); !val {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
