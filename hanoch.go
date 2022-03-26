package leetcode

//https://leetcode-cn.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

//https://leetcode-cn.com/problems/add-two-numbers/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

//https://leetcode-cn.com/problems/baseball-game/
func CalPoints(ops []string) int {
	return calPoints(ops)
}

//https://leetcode-cn.com/problems/3sum/
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
