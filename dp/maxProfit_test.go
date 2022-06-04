package dp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(maxProfit([]int{7, 1, 5, 3, 6, 4}), ShouldEqual, 5)
		})
		Convey("2", func() {
			So(maxProfit([]int{7, 6, 5, 4, 3, 2, 1}), ShouldEqual, 0)
		})
	})
}
