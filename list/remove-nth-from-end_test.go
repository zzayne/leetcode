package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	Convey("test1", t, func() {
		node := getNodeList([]int{1, 2})
		node = removeNthFromEnd(node, 1)
		So(getListNums(node), ShouldResemble, []int{1})
	})
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
