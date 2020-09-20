package leetcode

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Length int
	Head   *ListNode
}

func (ll *List) Add(Val int) bool {
	node := &ListNode{
		Val: Val,
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
func TestIsPalindrome(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		l1 := &List{}
		l1.Add(1)
		l1.Add(2)

		want := false
		got := isPalindrome(l1.Head)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		l2 := &List{}
		l2.Add(1)
		l2.Add(2)
		l2.Add(2)
		l2.Add(1)

		want := true
		got := isPalindrome(l2.Head)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
