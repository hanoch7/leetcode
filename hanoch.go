package leetcode

// https://leetcode-cn.com/problems/two-sum/
func TwoSum(nums []int, target int) []int {
	return twoSum(nums, target)
}

// https://leetcode-cn.com/problems/add-two-numbers/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

// https://leetcode-cn.com/problems/baseball-game/
func CalPoints(ops []string) int {
	return calPoints(ops)
}

// https://leetcode-cn.com/problems/3sum/
func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

// https://leetcode-cn.com/problems/generate-parentheses/
func GenerateParenthesis(n int) []string {
	return generateParenthesis(n)
}

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func ReverseKGroup(head *ListNode, k int) *ListNode {
	return reverseKGroup(head, k)
}

// https://leetcode-cn.com/problems/find-missing-observations/
func MissingRolls(rolls []int, mean int, n int) []int {
	return missingRolls(rolls, mean, n)
}

// https://leetcode-cn.com/problems/trapping-rain-water/
func Trap(height []int) int {
	return trap(height)
}

// https://leetcode-cn.com/problems/jump-game-ii/
func Jump(nums []int) int {
	return jump(nums)
}

// https://leetcode-cn.com/problems/permutations/
func Permute(nums []int) [][]int {
	return permute(nums)
}

// https://leetcode-cn.com/problems/binary-number-with-alternating-bits/
func HasAlternatingBits(n int) bool {
	return hasAlternatingBits(n)
}

// https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam/
func MaxConsecutiveAnswers(answerKey string, k int) int {
	return maxConsecutiveAnswers(answerKey, k)
}

// https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees/
func GetAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	return getAllElements(root1, root2)
}

// https://leetcode-cn.com/problems/subarray-product-less-than-k/
func NumSubarrayProductLessThanK(nums []int, k int) int {
	return numSubarrayProductLessThanK(nums, k)
}

// https://leetcode.cn/problems/di-string-match/
func DiStringMatch(s string) []int {
	return diStringMatch(s)
}

// https://leetcode.cn/problems/height-checker/
func HeightChecker(heights []int) int {
	return heightChecker(heights)
}

// https://leetcode.cn/problems/jump-game/
func CanJump(nums []int) bool {
	return canJump(nums)
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
func DeleteDuplicates(head *ListNode) *ListNode {
	return deleteDuplicates(head)
}

//https://leetcode.cn/problems/climbing-stairs/
func ClimbStairs(n int) int {
	return climbStairs(n)
}

//https://leetcode.cn/problems/path-sum/
func HasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSum(root, targetSum)
}

//https://leetcode.cn/problems/fibonacci-number/
func Fib(n int) int {
	return fib(n)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
