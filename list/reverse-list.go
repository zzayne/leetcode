package list

// https://leetcode.cn/problems/reverse-linked-list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var lst []*ListNode
	for node := head; node != nil; node = node.Next {
		lst = append(lst, node)
	}

	length := len(lst)

	var cur *ListNode
	for i := 0; i < len(lst); i++ {
		node := ListNode{
			Val:  lst[i].Val,
			Next: cur,
		}
		if i == length-1 {
			head = &node
		}
		cur = &node
	}
	return cur
}
