package leetcode

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node != nil {
			ans = append(ans, node.Val)
			preorder(node.Left)
			preorder(node.Right)
		}
	}

	preorder(root)

	return ans
}
