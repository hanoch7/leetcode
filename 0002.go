package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	tmp := ans
	advance := 0

	for l1 != nil && l2 != nil {
		tmp.Next = &ListNode{}
		tmp = tmp.Next
		val := l1.Val + l2.Val + advance
		tmp.Val = val % 10
		advance = val / 10
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		tmp.Next = &ListNode{}
		tmp = tmp.Next
		val := l1.Val + advance
		tmp.Val = val % 10
		advance = val / 10
		l1 = l1.Next
	}

	for l2 != nil {
		tmp.Next = &ListNode{}
		tmp = tmp.Next
		val := l2.Val + advance
		tmp.Val = val % 10
		advance = val / 10
		l2 = l2.Next
	}

	if advance != 0 {
		tmp.Next = &ListNode{}
		tmp = tmp.Next
		tmp.Val = advance
	}

	return ans.Next
}
