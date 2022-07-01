package leetcode

func canJump(nums []int) bool {
	canReach := 0
	nowStep := 0
	lastStep := len(nums) - 1
	for nowStep <= canReach {
		if nowStep+nums[nowStep] >= canReach {
			canReach = nowStep + nums[nowStep]
			if canReach >= lastStep {
				return true
			}
		}
		nowStep += 1
	}

	return false
}
