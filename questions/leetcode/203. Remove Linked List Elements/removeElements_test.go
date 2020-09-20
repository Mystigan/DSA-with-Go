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
func TestRemoveElements(t *testing.T) {
	l1 := &List{}
	l1.Add(1)
	l1.Add(2)
	l1.Add(6)
	l1.Add(3)
	l1.Add(4)
	l1.Add(5)
	l1.Add(6)
	want := []int{1, 2, 3, 4, 5}
	ans := removeElements(l1.Head, 6)
	got := []int{ans.Val}
	for ans.Next != nil {
		ans = ans.Next
		got = append(got, ans.Val)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
