package list

func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow

}
