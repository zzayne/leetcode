package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GetTree(nums []*int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var nodes []*TreeNode
	for i := 0; i < len(nums); i++ {
		nodes = append(nodes, getNode(nums, i))
	}

	for i := 0; i < len(nodes); i++ {
		if nodes[i] == nil {
			continue
		}
		idx := i * 2
		if idx+1 < len(nodes) {
			nodes[i].Left = nodes[idx+1]
		}
		if idx+2 < len(nodes) {
			nodes[i].Right = nodes[idx+2]
		}
	}
	return nodes[0]
}

func getNode(nums []*int, idx int) *TreeNode {
	if idx > len(nums)-1 {
		return nil
	}

	if nums[idx] == nil {
		return nil
	}
	return &TreeNode{
		Val:   *nums[idx],
		Left:  nil,
		Right: nil,
	}
}
