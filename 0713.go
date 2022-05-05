package leetcode

func numSubarrayProductLessThanK(nums []int, k int) int {
	ans, i, j, tmp := 0, 0, 0, 1
	for j < len(nums) {
		tmp *= nums[j]
		j++
		for tmp >= k {
			tmp /= nums[i]
			i++
		}
		ans += j - i
	}
	return ans
}
