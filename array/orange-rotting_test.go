package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_orangesRotting(t *testing.T) {
	Convey("test1", t, func() {
		grid := [][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}
		So(orangesRotting(grid), ShouldEqual, 4)
		//So(, ShouldEqual, -1)
	})
}
