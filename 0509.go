package leetcode

func fib(n int) int {
	first := 0
	if n == 0 {
		return first
	}

	second := 1
	if n == 1 {
		return second
	}

	now := 0
	for i := 2; i <= n; i++ {
		now = first + second
		first = second
		second = now
	}
	return now
}
