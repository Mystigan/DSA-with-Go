package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	aLen, bLen := 1, 1
	currA, currB := headA, headB
	for currA.Next != nil {
		currA = currA.Next
		aLen++
	}

	for currB.Next != nil {
		currB = currB.Next
		bLen++
	}
	p1, p2, diff := headA, headB, aLen-bLen
	if bLen > aLen {
		p1 = headB
		p2 = headA
		diff = bLen - aLen
	}

	for diff > 0 {
		p1 = p1.Next
		diff--
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
