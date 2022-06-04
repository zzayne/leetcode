package dp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), ShouldEqual, 6)
		})
		Convey("2", func() {
			So(maxSubArray([]int{5, 4, -1, 7, 8}), ShouldEqual, 23)
		})
	})
}
