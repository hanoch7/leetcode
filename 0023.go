package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeLists(lists, 0, len(lists)-1)
}

func mergeLists(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left == right-1 {
		return mergeTwoLists(lists[left], lists[right])
	}
	mid := (left + right) / 2
	return mergeTwoLists(mergeLists(lists, left, mid), mergeLists(lists, mid+1, right))
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	now := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			now.Next = l1
			l1 = l1.Next
			now = now.Next
		} else {
			now.Next = l2
			l2 = l2.Next
			now = now.Next
		}
	}
	if l1 != nil {
		now.Next = l1
	} else {
		now.Next = l2
	}
	return head.Next
}
