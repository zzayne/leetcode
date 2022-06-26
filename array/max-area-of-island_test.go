package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_maxAreaOfIsland(t *testing.T) {
	Convey("test1", t, func() {
		grid := [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		}
		So(maxAreaOfIsland(grid), ShouldEqual, 6)
		//So(, ShouldEqual, -1)
	})
	Convey("test2", t, func() {
		grid := [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
		}
		So(maxAreaOfIsland(grid), ShouldEqual, 0)
		//So(, ShouldEqual, -1)
	})
}
