package tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxInt := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	return maxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
