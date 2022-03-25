package leetcode

//https://leetcode-cn.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

//https://leetcode-cn.com/problems/add-two-numbers/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
