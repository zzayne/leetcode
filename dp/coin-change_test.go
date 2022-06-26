package dp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_coinChange(t *testing.T) {
	Convey("test1", t, func() {
		//Convey("1", func() {
		//	So(coinChange([]int{1, 2, 5}, 11), ShouldEqual, 3)
		//})
		Convey("2", func() {
			So(coinChange([]int{2}, 3), ShouldEqual, -1)
		})
	})
}
