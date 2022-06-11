package math

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(isPalindrome(121), ShouldEqual, true)
		})
		Convey("2", func() {
			So(isPalindrome(-121), ShouldEqual, false)
		})
		Convey("3", func() {
			So(isPalindrome(10), ShouldEqual, false)
		})
	})
}
