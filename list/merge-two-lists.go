package list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	newHead := new(ListNode)
	if list1.Val < list2.Val {
		newHead = mergeTwoLists(list1.Next, list2)
		list1.Next = newHead
		newHead = list1
	} else {
		newHead = mergeTwoLists(list1, list2.Next)
		list2.Next = newHead
		newHead = list2
	}
	return newHead
}
