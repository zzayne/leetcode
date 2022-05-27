package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	Convey("isPalindrome", t, func() {
		Convey("1", func() {
			So(isPalindrome("A man, a plan, a canal: Panama"), ShouldEqual, true)
		})
		Convey("2", func() {
			So(isPalindrome("race a car"), ShouldEqual, false)
		})
		Convey("3", func() {
			So(isPalindrome("rac "), ShouldEqual, false)
		})
		Convey("4", func() {
			So(isPalindrome("0P"), ShouldEqual, false)
		})
		Convey("5", func() {
			So(isPalindrome(" rr   "), ShouldEqual, true)
		})
	})
}
