package leetcode

var perMuteAns [][]int

func permute(nums []int) [][]int {
	perMuteAns = [][]int{}
	getPermute(nums, 0, len(nums))
	return perMuteAns
}

func getPermute(nums []int, changed int, n int) {
	if changed == n {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		perMuteAns = append(perMuteAns, tmp)
		return
	}

	for i := changed; i < n; i++ {
		nums[changed], nums[i] = nums[i], nums[changed]
		getPermute(nums, changed+1, n)
		nums[changed], nums[i] = nums[i], nums[changed]
	}
}
