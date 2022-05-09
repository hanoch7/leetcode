package leetcode

func diStringMatch(s string) []int {
	i, j := 0, len(s)
	ans := []int{}
	for _, k := range s {
		if k == 'I' {
			ans = append(ans, i)
			i++
		} else {
			ans = append(ans, j)
			j--
		}
	}
	ans = append(ans, i)
	return ans
}
