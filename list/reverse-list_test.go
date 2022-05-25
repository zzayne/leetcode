package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reverseList(t *testing.T) {
	Convey("test1", t, func() {
		node := getNodeList([]int{1, 2, 3, 4, 5})
		node = reverseList(node)
		So(getListNums(node), ShouldResemble, []int{5, 4, 3, 2, 1})
	})
	Convey("test2", t, func() {
		node := getNodeList([]int{1})
		node = reverseList(node)
		So(getListNums(node), ShouldResemble, []int{1})
	})
	Convey("test3", t, func() {
		node := getNodeList([]int{})
		node = reverseList(node)
		So(len(getListNums(node)), ShouldEqual, 0)
	})
}
