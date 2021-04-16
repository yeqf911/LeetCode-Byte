package structure

type ListNode struct {
	Val  int
	Next *ListNode
}

func Array2Linklist(nums []int) *ListNode {
	head := &ListNode{}
	p := head
	for _, n := range nums {
		p.Next = &ListNode{}
		p = p.Next
		p.Val = n
	}

	return head.Next
}
