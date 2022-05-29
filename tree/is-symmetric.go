package tree

//https://leetcode.cn/problems/symmetric-tree/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var check func(*TreeNode, *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}
	return check(root.Left, root.Right)
}
