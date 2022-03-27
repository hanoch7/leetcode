package leetcode

func jump(nums []int) int {
	maxReach, curReach, i, step := 0, 0, 0, 0
	for i < len(nums)-1 {
		maxReach = max(maxReach, nums[i]+i)
		if curReach == i {
			step++
			curReach = maxReach
		}
		i++
	}
	return step
}
