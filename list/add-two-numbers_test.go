package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			res := addTwoNumbers(getNodeList([]int{2, 4, 3}), getNodeList([]int{5, 6, 4}))
			So(getListNums(res), ShouldResemble, []int{7, 0, 8})
		})
		Convey("2", func() {
			res := addTwoNumbers(getNodeList([]int{0}), getNodeList([]int{0}))
			So(getListNums(res), ShouldResemble, []int{0})
		})
		Convey("3", func() {
			res := addTwoNumbers(getNodeList([]int{9, 9, 9, 9, 9, 9, 9}), getNodeList([]int{9, 9, 9, 9}))
			So(getListNums(res), ShouldResemble, []int{8, 9, 9, 9, 0, 0, 0, 1})
		})
		Convey("4", func() {
			res := addTwoNumbers(getNodeList([]int{1, 0, 1}), getNodeList([]int{5}))
			So(getListNums(res), ShouldResemble, []int{6, 0, 1})
		})
		Convey("5", func() {
			res := addTwoNumbers(getNodeList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}), getNodeList([]int{5, 6, 4}))
			So(getListNums(res), ShouldResemble, []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
		})
		Convey("6", func() {
			res := addTwoNumbers(getNodeList([]int{9, 9}), getNodeList([]int{1}))
			So(getListNums(res), ShouldResemble, []int{0, 0, 1})
		})
	})
}
