package leetcode

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	for _, nums := range matrix {
		i := sort.SearchInts(nums, target)
		if i < len(nums) && nums[i] == target {
			return true
		}
	}
	return false
}
