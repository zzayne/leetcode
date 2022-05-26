package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			node := getNodeList([]int{1, 2, 1})
			So(isPalindrome(node), ShouldEqual, true)
		})
		Convey("2", func() {
			node := getNodeList([]int{1, 2, 2, 1})
			So(isPalindrome(node), ShouldEqual, true)
		})

		Convey("3", func() {
			node := getNodeList([]int{1, 2, 3})
			So(isPalindrome(node), ShouldEqual, false)
		})
		Convey("4", func() {
			node := getNodeList([]int{1, 2, 3, 2, 1})
			So(isPalindrome(node), ShouldEqual, true)
		})
		Convey("5", func() {
			node := getNodeList([]int{1, 2})
			So(isPalindrome(node), ShouldEqual, false)
		})
		Convey("6", func() {
			node := getNodeList([]int{1})
			So(isPalindrome(node), ShouldEqual, true)
		})
	})
}
