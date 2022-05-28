package tree

//https://leetcode.cn/problems/validate-binary-search-tree/
// 考察点：中序遍历
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var travel func(*TreeNode) bool
	travel = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !travel(node.Left) {
			return false
		}
		if prev != nil && prev.Val >= node.Val {
			return false
		}

		prev = node
		if !travel(node.Right) {
			return false
		}
		return true
	}
	return travel(root)

}
