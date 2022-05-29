package tree

// https://leetcode.cn/problems/binary-tree-level-order-traversal/
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var travel func(*TreeNode, *TreeNode, int)
	res = append(res, []int{root.Val})
	travel = func(left *TreeNode, right *TreeNode, i int) {
		if left == nil && right == nil {
			return
		}

		for len(res)-1 < i {
			res = append(res, []int{})
		}
		next := i + 1
		if left != nil {
			res[i] = append(res[i], left.Val)
			travel(left.Left, left.Right, next)
		}
		if right != nil {
			res[i] = append(res[i], right.Val)
			travel(right.Left, right.Right, next)
		}
	}

	travel(root.Left, root.Right, 1)
	return res
}
