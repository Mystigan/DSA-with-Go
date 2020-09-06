package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// vals := []int{}
	// if head==nil || head.Next==nil {
	//     return head
	// }
	// current:= head
	// vals = append(vals, current.Val)
	// for current.Next!=nil {
	//     current = current.Next
	//     vals = append(vals, current.Val)
	// }
	// current = head
	// for i:=len(vals)-1; i>=0; i-- {
	//     current.Val = vals[i]
	//     if current.Next!=nil{current = current.Next}
	// }
	// return head

	//alternate solution
	var prev *ListNode
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
