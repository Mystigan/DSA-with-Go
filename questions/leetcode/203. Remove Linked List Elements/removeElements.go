package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		if head.Val == val {
			return nil
		}
		return head
	}
	var temp *ListNode
	for head.Val == val && head.Next != nil {
		temp = head.Next
		head.Next = nil
		head = temp
	}
	curr := head
	prev := head

	for curr != nil {
		if curr.Val == val {
			if prev != curr {
				prev.Next = curr.Next
				curr = curr.Next
				continue
			} else {
				if curr.Next == nil {
					return nil
				}
				prev = curr.Next
				curr.Next = nil
				curr = prev
				continue
			}
		}
		prev = curr
		curr = curr.Next
	}
	return head
}
