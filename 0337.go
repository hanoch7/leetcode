package leetcode

func rob3(root *TreeNode) int {
	v := dfsRob(root)
	return max(v[0], v[1])
}

func dfsRob(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	l := dfsRob(node.Left)
	r := dfsRob(node.Right)
	se := node.Val + l[1] + r[1]
	notSe := max(l[0], l[1]) + max(r[0], r[1])
	return []int{se, notSe}
}
