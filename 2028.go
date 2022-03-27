package leetcode

func missingRolls(rolls []int, mean int, n int) []int {
	ans := []int{}

	sumM := 0
	for _, roll := range rolls {
		sumM += roll
	}

	sumN := mean*(len(rolls)+n) - sumM
	if sumN > 6*n || sumN < 1*n {
		return ans
	}
	for i := 0; i < n; i++ {
		if i < sumN%n {
			ans = append(ans, sumN/n+1)
		} else {
			ans = append(ans, sumN/n)
		}

	}

	return ans
}
