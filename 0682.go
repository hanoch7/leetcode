package leetcode

import "strconv"

func calPoints(ops []string) int {
	a := []int{}
	for _, op := range ops {
		switch op {
		case "+":
			a = append(a, a[len(a)-1]+a[len(a)-2])
		case "D":
			a = append(a, a[len(a)-1]*2)
		case "C":
			a = a[:len(a)-1]
		default:
			i, _ := strconv.Atoi(op)
			a = append(a, i)
		}
	}
	ans := 0
	for _, i := range a {
		ans += i
	}
	return ans
}
