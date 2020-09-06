package leetcode

import (
	"strconv"
	"strings"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	current := head
	var numstring strings.Builder
	numstring.WriteString(strconv.Itoa(current.Val))
	for current.Next != nil {
		current = current.Next
		numstring.WriteString(strconv.Itoa(current.Val))
	}

	num, _ := strconv.ParseInt(numstring.String(), 2, 32)
	return int(num)
}
