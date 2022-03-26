package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for k := range nums {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		for i, j := k+1, len(nums)-1; i < j; {
			if i > k+1 && nums[i] == nums[i-1] {
				i++
				continue
			}
			if j < len(nums)-1 && nums[j] == nums[j+1] {
				j--
				continue
			}
			if nums[k]+nums[i]+nums[j] == 0 {
				ans = append(ans, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				continue
			}
			if nums[k]+nums[i]+nums[j] < 0 {
				i++
				continue
			}
			if nums[k]+nums[i]+nums[j] > 0 {
				j--
				continue
			}
		}
	}
	return ans
}
