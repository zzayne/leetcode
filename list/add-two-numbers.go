package list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	var res, prev *ListNode

	addNode := func(val int) {
		node := ListNode{
			Val: val,
		}
		if res == nil {
			res = &node
			prev = &node

		} else {
			prev.Next = &node
			prev = &node
		}
	}
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if flag > 0 {
			sum += flag
		}
		if sum >= 10 {
			flag = 1
			sum = sum - 10
		} else {
			flag = 0
		}
		addNode(sum)
	}
	if flag > 0 {
		addNode(flag)
	}
	return res
}
