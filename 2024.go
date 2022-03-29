package leetcode

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveChar(answerKey, k, 'T'), maxConsecutiveChar(answerKey, k, 'F'))
}

func maxConsecutiveChar(answerKey string, k int, ch byte) (ans int) {
	l, sum := 0, 0
	for r := range answerKey {
		if answerKey[r] != ch {
			sum++
		}
		for sum > k {
			if answerKey[l] != ch {
				sum--
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return
}
