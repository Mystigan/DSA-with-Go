package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//// Time: O(n), space: O(n)
// func isPalindrome(head *ListNode) bool {
//     if head==nil ||  head.Next==nil{
//         return true
//     }
//     nodes := []int{head.Val}
//     for head.Next!=nil {
//         head = head.Next
//         nodes = append(nodes, head.Val)
//     }
//     n := len(nodes)
//     for i:=0; i<n/2; i++ {
//         if nodes[i] != nodes[n-1-i] {
//             return false
//         }
//     }
//     return true
// }

// Time: O(n), space: O(1)
func isPalindrome(head *ListNode) bool {
	// base case
	if head == nil || head.Next == nil {
		return true
	}

	// 1. find the half
	var prev *ListNode = nil
	var slow *ListNode = head
	var fast *ListNode = head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// cut half
	prev.Next = nil

	// 2. set the second half and reverse the first half
	if fast != nil {
		slow = slow.Next
	}
	var head2 *ListNode = slow
	head = reverse(head)

	// 3. compare two linked lists
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var curr *ListNode = head

	for curr != nil {
		var post *ListNode = curr.Next
		curr.Next = prev
		prev = curr
		curr = post
	}

	return prev
}
