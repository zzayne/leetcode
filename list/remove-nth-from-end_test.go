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
