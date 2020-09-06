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

func (ll *List) Add(val int) {
	node := &ListNode{
		Val: val,
	}
	if ll.Head == nil {
		ll.Head = node
		ll.Length++
	} else {
		// Adding to the end of the list
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
		ll.Length++

		// // Adding to the start of the list
		// node.Next = ll.Head
		// ll.Head = node
		// ll.Length++
	}
}

func (ll *List) Find(val int) (*ListNode, int) {
	// Not checking for empty list as we are are guaranteed to have a list with required value.
	current, pos := ll.Head, 0
	for current.Next != nil {
		if current.Val == val {
			return current, pos
		}
		current = current.Next
		pos++
	}
	return nil, 0
}

func (ll *List) FindAt(pos int) int {
	current := ll.Head
	for pos > 0 {
		current = current.Next
		pos--
	}
	return current.Val
}

// Note: The tests are written in such a way to avoid any modifications to the solutions function as given by Leetcode.
func TestDeleteNode(t *testing.T) {
	ll := &List{}
	ll.Add(4)
	ll.Add(5)
	ll.Add(1)
	ll.Add(9)

	want := 1
	node, pos := ll.Find(5)
	deleteNode(node)
	got := ll.FindAt(pos)

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

}
