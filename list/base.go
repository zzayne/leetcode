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

func getListNums(head *ListNode) (out []int) {
	node := head
	out = append(out, node.Val)
	for node.Next != nil {
		node = node.Next
		out = append(out, node.Val)
	}
	return out
}
