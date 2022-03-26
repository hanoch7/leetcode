package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	hear := &ListNode{
		Next: head,
	}
	tmp1, tmp2 := hear, hear.Next
	for tmp2 != nil {
		tmp4 := tmp2
		flag := true
		for i := 0; i < k-1; i++ {
			tmp4 = tmp4.Next
			if tmp4 == nil {
				flag = false
				tmp2 = nil
				break
			}
		}
		if flag {
			for i := 0; i < k-1; i++ {
				tmp3 := tmp2.Next
				tmp2.Next = tmp3.Next
				tmp3.Next = tmp1.Next
				tmp1.Next = tmp3
			}
			tmp1 = tmp2
			tmp2 = tmp1.Next
		}
	}
	return hear.Next
}
