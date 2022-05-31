package math

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isPowerOfThree(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(isPowerOfThree(27), ShouldEqual, true)
		})
		Convey("2", func() {
			So(isPowerOfThree(0), ShouldEqual, false)
		})
		Convey("3", func() {
			So(isPowerOfThree(9), ShouldEqual, true)
		})
		Convey("4", func() {
			So(isPowerOfThree(45), ShouldEqual, false)
		})
		Convey("5", func() {
			So(isPowerOfThree(1), ShouldEqual, true)
		})

	})
}
