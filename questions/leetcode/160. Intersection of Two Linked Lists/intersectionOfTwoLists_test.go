package leetcode

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Length int
	Head   *ListNode
}

func (ll *List) Add(data int) bool {
	node := &ListNode{
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

func TestGetIntersectionNode(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		list1 := &List{}
		list1.Add(4)
		list1.Add(1)
		list1.Add(8)
		list1.Add(4)
		list1.Add(5)

		list2 := &List{}
		list2.Add(5)
		list2.Add(6)
		list2.Add(1)
		list2.Head.Next.Next.Next = list1.Head.Next.Next
		want := []int{8, 4, 5}
		val := getIntersectionNode(list1.Head, list2.Head)
		got := []int{val.Val}
		for val.Next != nil {
			val = val.Next
			got = append(got, val.Val)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v ", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		list1 := &List{}
		list1.Add(1)
		list2 := &List{}
		list2.Add(2)
		list2.Add(3)
		var want *ListNode = nil
		got := getIntersectionNode(list1.Head, list2.Head)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
