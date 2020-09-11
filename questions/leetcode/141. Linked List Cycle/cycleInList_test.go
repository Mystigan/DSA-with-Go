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

func TestHasCycle(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		l1 := &List{}
		l1.Add(3)
		l1.Add(2)
		l1.Add(0)
		l1.Add(-4)
		var posNode *ListNode
		curr := l1.Head
		for curr.Next != nil {
			curr = curr.Next
			if curr.Val == 2 {
				posNode = curr
			}
		}
		curr.Next = posNode
		want := true
		got := hasCycle(l1.Head)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test two", func(t *testing.T) {
		l1 := &List{}
		l1.Add(1)
		l1.Add(2)
		l1.Head.Next.Next = l1.Head
		want := true
		got := hasCycle(l1.Head)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("test three", func(t *testing.T) {
		l1 := &List{}
		l1.Add(1)
		want := false
		got := hasCycle(l1.Head)
		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
