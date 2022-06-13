package list

func reversePrint(head *ListNode) []int {
	var res []int

	for ; head != nil; head = head.Next {
		res = append(res, head.Val)
	}

	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
