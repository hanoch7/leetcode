package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	first := nums[0]
	second := max(nums[0], nums[1])
	now := max(first+nums[2], second)

	for i := 2; i < len(nums); i++ {
		now = max(first+nums[i], second)
		first = second
		second = now
	}
	return now
}
