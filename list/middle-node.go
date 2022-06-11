package list

func middleNode(head *ListNode) *ListNode {
	count := 0

	for node := head; node != nil; node = node.Next {
		count++
	}
	mid := count / 2

	for i := 0; i < mid; i++ {
		head = head.Next
	}
	return head
}
