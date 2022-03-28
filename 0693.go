package leetcode

func hasAlternatingBits(n int) bool {
	pre := -1
	for n != 0 {
		cur := n % 2
		if cur == pre {
			return false
		}
		pre = cur
		n /= 2
	}
	return true
}
