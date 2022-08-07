package leetcode

func climbStairs(n int) int {
	first := 1
	if n == 1 {
		return first
	}

	second := 2
	if n == 2 {
		return second
	}

	now := 0
	for i := 3; i <= n; i++ {
		now = first + second
		first = second
		second = now
	}

	return now
}
