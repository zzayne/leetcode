package tree

import "container/list"

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var res [][]*Node
	queue := list.New()
	queue.PushBack(root)

	var tmpArr []*Node
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node)
		}
		res = append(res, tmpArr)
		tmpArr = []*Node{}
	}
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i])-1; j++ {
			res[i][j].Next = res[i][j+1]
		}
	}
	return root
}
