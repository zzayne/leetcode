package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			node := mergeTwoLists(getNodeList([]int{1, 2, 4}), getNodeList([]int{1, 3, 4}))
			So(getListNums(node), ShouldResemble, []int{1, 1, 2, 3, 4, 4})
		})
		Convey("2", func() {
			node := mergeTwoLists(nil, nil)
			So(len(getListNums(node)), ShouldEqual, 0)
		})
		Convey("3", func() {
			node := mergeTwoLists(nil, getNodeList([]int{1, 3, 4}))
			So(getListNums(node), ShouldResemble, []int{1, 3, 4})
		})
		Convey("4", func() {
			node := mergeTwoLists(getNodeList([]int{1, 3, 4}), nil)
			So(getListNums(node), ShouldResemble, []int{1, 3, 4})
		})
	})
}
