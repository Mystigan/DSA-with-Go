package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func middleNode(head *ListNode) *ListNode {
//     length, current := 1, head
//     for ; current.Next!=nil; current=current.Next {
//         length++
//     }
//     length = length/2 + 1
//     for length>1 {
//         head = head.Next
//         length--
//     }
//     return head
// }

//Two pointer based solution
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast.Next != nil && fast.Next.Next == nil {
		return slow.Next
	}
	return slow
}
