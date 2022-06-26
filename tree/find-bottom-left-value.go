package tree

import "container/list"

func findBottomLeftValue(root *TreeNode) int {
	var res int
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		n := queue.Len()

		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}
