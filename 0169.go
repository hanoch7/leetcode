package leetcode

func majorityElement(nums []int) int {
	now := nums[0]
	count := 0
	for _, num := range nums {
		if num == now {
			count += 1
			continue
		}
		count -= 1
		if count == 0 {
			now = num
			count = 1
			continue
		}
	}
	return now
}
