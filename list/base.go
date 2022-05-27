package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func getNodeList(nums []int) *ListNode {
	var head *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		node := ListNode{
			Val:  nums[i],
			Next: head,
		}
		head = &node
	}

	return head
}

func getCycleNodeList(nums []int, post int) *ListNode {
	var head *ListNode

	var nodes []*ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		node := ListNode{
			Val:  nums[i],
			Next: head,
		}
		head = &node
		nodes = append(nodes, &node)
	}

	if len(nodes) > 0 && len(nodes) > post {
		nodes[0].Next = nodes[len(nodes)-post-1]
	}
	return head
}

func getListNums(head *ListNode) (out []int) {
	if head == nil {
		return nil
	}
	node := head
	out = append(out, node.Val)
	for node.Next != nil {
		node = node.Next
		out = append(out, node.Val)
	}
	return out
}
