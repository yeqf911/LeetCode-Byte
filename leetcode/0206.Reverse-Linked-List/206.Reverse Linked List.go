package leetcode

import "LeetCode-Byte/structure"

type ListNode = structure.ListNode

func reverseList(head *ListNode) *ListNode {
	ptr := head
	var pre *ListNode
	for ptr != nil {
		next := ptr.Next
		ptr.Next = pre
		pre = ptr
		ptr = next
	}
	return pre
}