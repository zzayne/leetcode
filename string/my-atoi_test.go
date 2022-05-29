package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(myAtoi(" +-09"), ShouldEqual, 0)
		})
		Convey("2", func() {
			So(myAtoi("42"), ShouldEqual, 42)
		})
		Convey("3", func() {
			So(myAtoi("   -42"), ShouldEqual, -42)
		})
		Convey("4", func() {
			So(myAtoi("4193 with words"), ShouldEqual, 4193)
		})
		Convey("5", func() {
			So(myAtoi("words and 987"), ShouldEqual, 0)
		})

		Convey("6", func() {

			So(myAtoi("-91283472332"), ShouldEqual, -2147483648)
		})

	})
}
