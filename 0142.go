package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head.Next
	if slow.Next == head {
		return head
	}

	fast := slow.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	fast = head

	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
