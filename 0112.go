package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return hPS(root, targetSum)
}

func hPS(node *TreeNode, targetSum int) bool {
	if node == nil {
		return targetSum == 0
	}

	if node.Left == nil {
		return hPS(node.Right, targetSum-node.Val)
	}

	if node.Right == nil {
		return hPS(node.Left, targetSum-node.Val)
	}

	return hPS(node.Left, targetSum-node.Val) || hPS(node.Right, targetSum-node.Val)
}
