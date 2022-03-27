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

//https://leetcode-cn.com/problems/generate-parentheses/
func GenerateParenthesis(n int) []string {
	return generateParenthesis(n)
}

//https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	return reverseKGroup(head, k)
}

//https://leetcode-cn.com/problems/find-missing-observations/
func MissingRolls(rolls []int, mean int, n int) []int {
	return missingRolls(rolls, mean, n)
}

//https://leetcode-cn.com/problems/trapping-rain-water/
func Trap(height []int) int {
	return trap(height)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
