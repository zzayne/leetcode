package dp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(climbStairs(2), ShouldEqual, 2)
		})
		Convey("2", func() {
			So(climbStairs(3), ShouldEqual, 3)
		})
	})
}
