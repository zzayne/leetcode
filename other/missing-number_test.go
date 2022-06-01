package other

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_missingNumber(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(missingNumber([]int{3, 0, 1}), ShouldEqual, 2)
		})

		Convey("2", func() {
			So(missingNumber([]int{0, 1}), ShouldEqual, 2)
		})

		Convey("3", func() {
			So(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), ShouldEqual, 8)
		})

		Convey("4", func() {
			So(missingNumber([]int{0}), ShouldEqual, 1)
		})

		//So(, ShouldEqual, -1)
	})
}
