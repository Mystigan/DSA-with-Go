package leetcode

import (
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

func (ll *List) Add(val int) {
	node := &ListNode{
		Val: val,
	}
	if ll.Head == nil {
		ll.Head = node
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

func TestGetDecimalValue(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		nums := []int{1, 0, 1}
		ll := &List{}
		for _, v := range nums {
			ll.Add(v)
		}
		want := 5
		got := getDecimalValue(ll.Head)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		nums := []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}
		ll := &List{}
		for _, v := range nums {
			ll.Add(v)
		}
		want := 18880
		got := getDecimalValue(ll.Head)

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
