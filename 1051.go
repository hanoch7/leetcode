package leetcode

func heightChecker(heights []int) int {
	tmp := [101]int{}
	for _, h := range heights {
		tmp[h]++
	}
	ans := 0
	idx := 0
	for i, t := range tmp {
		for t > 0 {
			if heights[idx] != i {
				ans++
			}
			idx++
			t--
		}
	}
	return ans
}
