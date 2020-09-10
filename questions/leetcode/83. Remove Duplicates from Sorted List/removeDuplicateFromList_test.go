package leetcode

import (
	"reflect"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func TestDeleteDuplicates(t *testing.T) {
	t.Run("test one", func(t *testing.T) {
		l1 := &List{}
		l1.Add(1)
		l1.Add(1)
		l1.Add(2)

		want := []int{1, 2}
		ans := deleteDuplicates(l1.Head)
		got := []int{}
		got = append(got, ans.Val)
		for ans.Next != nil {
			ans = ans.Next
			got = append(got, ans.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("test two", func(t *testing.T) {
		l2 := &List{}
		l2.Add(1)
		l2.Add(1)
		l2.Add(2)
		l2.Add(3)
		l2.Add(3)

		want := []int{1, 2, 3}
		ans := deleteDuplicates(l2.Head)
		got := []int{}
		got = append(got, ans.Val)
		for ans.Next != nil {
			ans = ans.Next
			got = append(got, ans.Val)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
