package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_middleNode(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			res := middleNode(getNodeList([]int{1, 2, 3, 4, 5, 6}))
			So(getListNums(res), ShouldResemble, []int{4, 5, 6})
		})
		Convey("2", func() {
			res := middleNode(getNodeList([]int{1}))
			So(getListNums(res), ShouldResemble, []int{1})
		})
		Convey("3", func() {
			res := middleNode(getNodeList([]int{1, 1}))
			So(res, ShouldResemble, getNodeList([]int{1}))
		})
	})
}
