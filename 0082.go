package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	hear := &ListNode{
		Next: head,
	}
	zero := hear
	first := head
	for first != nil && first.Next != nil {
		second := first.Next
		if first.Val == second.Val {
			for second.Next != nil && second.Next.Val == second.Val {
				second = second.Next
			}
			zero.Next = second.Next
			first = zero.Next
		} else {
			zero = first
			first = second
		}
	}

	return hear.Next
}
