package tree

import (
	"math"
)

//https://leetcode.cn/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	var lst []*TreeNode
	min := math.MinInt

	for len(lst) != 0 || root != nil {
		for root != nil {
			lst = append(lst, root)
			root = root.Left
		}
		root = lst[len(lst)-1]
		lst = lst[:len(lst)-1]
		if root.Val <= min {
			return false
		}
		min = root.Val
		root = root.Right
	}
	return true

}

// 考察点：中序遍历
func isValidBST1(root *TreeNode) bool {
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

// 递归遍历
func isValidBST2(root *TreeNode) bool {
	var travel func(node *TreeNode, min, max int) bool
	travel = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return travel(node.Left, min, node.Val) && travel(node.Right, node.Val, max)
	}
	return travel(root, math.MinInt, math.MaxInt)
}
