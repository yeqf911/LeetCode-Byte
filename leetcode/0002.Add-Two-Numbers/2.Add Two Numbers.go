package leetcode

import "LeetCode-Byte/structure"

type ListNode = structure.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	n1, n2, carry := 0, 0, 0
	p := head
	// carry != 0 是因为最后一位如果是进位的话，需要再遍历一次
	for l1 != nil || l2 != nil || carry != 0{
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}

		current := n1 + n2 + carry
		carry = current / 10
		p.Next = &ListNode{Val: current % 10}
		p = p.Next
	}
	return head.Next
}
