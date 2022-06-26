package tree

import (
	"container/list"
	"math"
)

func largestValues(root *TreeNode) []int {
	var res []int
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		n := queue.Len()

		max := math.MinInt
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, max)
	}
	return res
}
