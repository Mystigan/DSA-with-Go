package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if fast == slow || fast.Next == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
