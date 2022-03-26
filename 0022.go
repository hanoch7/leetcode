package leetcode

var ans []string

func generateParenthesis(n int) []string {
	ans = []string{}
	getParenthesis("", 0, 0, n)
	return ans
}

func getParenthesis(s string, left_count int, right_count int, n int) {
	if left_count == n && right_count == n {
		ans = append(ans, s)
		return
	}
	if left_count < right_count {
		return
	}
	if left_count > n {
		return
	}
	getParenthesis(s+"(", left_count+1, right_count, n)
	getParenthesis(s+")", left_count, right_count+1, n)
	return
}
