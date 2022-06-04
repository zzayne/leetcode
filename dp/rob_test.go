package dp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_rob(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(rob([]int{1, 2, 3, 1}), ShouldEqual, 4)
		})
		Convey("2", func() {
			So(rob([]int{2, 7, 9, 3, 1}), ShouldEqual, 12)
		})
		Convey("3", func() {
			So(rob([]int{1}), ShouldEqual, 1)
		})
		Convey("4", func() {
			So(rob([]int{1, 2}), ShouldEqual, 2)
		})

	})
}
