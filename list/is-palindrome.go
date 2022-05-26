package list

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	var lst []int
	for node := head; node != nil; node = node.Next {
		lst = append(lst, node.Val)
	}
	for i := 0; i < len(lst)/2; i++ {
		j := len(lst) - i - 1

		if lst[i] != lst[j] {
			return false
		}
	}
	return true
}
