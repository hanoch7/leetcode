package leetcode

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	if len(nums) == 3 {
		return max(nums[0], max(nums[1], nums[2]))
	}

	first := nums[0]
	second := max(nums[0], nums[1])
	now := max(first+nums[2], second)

	for i := 2; i < len(nums)-1; i++ {
		now = max(first+nums[i], second)
		first = second
		second = now
	}
	now_max := now

	first = nums[1]
	second = max(nums[1], nums[2])
	now = max(first+nums[3], second)
	for i := 3; i < len(nums); i++ {
		now = max(first+nums[i], second)
		first = second
		second = now
	}
	now_max = max(now, now_max)
	return now_max
}
