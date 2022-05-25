package list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var lst []*ListNode
	if head == nil {
		return nil
	}
	if head.Next == nil && n == 1 {
		return nil
	}
	node := head
	lst = append(lst, node)
	for node.Next != nil {
		if n == 1 && node.Next.Next == nil {
			node.Next = nil
			return head
		}
		node = node.Next
		lst = append(lst, node)
	}

	if len(lst) == n {
		return head.Next
	}

	node = lst[len(lst)-n]
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return head
}
