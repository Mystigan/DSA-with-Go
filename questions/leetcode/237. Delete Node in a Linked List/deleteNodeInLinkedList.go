package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// This approach shifts all elements to the left. Leaves no connections left to the last element.
// func deleteNode(node *ListNode) {
//     for node.Next!=nil {
//         node.Val = node.Next.Val
//         if node.Next.Next==nil{
//             node.Next = nil
//             break
//         }
//         node = node.Next
//     }
// }

//This approach only shift the next element to the left and then bypasses it. The right hand side link remaining on the next node apparently doesn't cause an issue.
//alternate solution
func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}
