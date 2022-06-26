package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_floodFill(t *testing.T) {
	Convey("test1", t, func() {
		image := [][]int{
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 0},
		}
		res := [][]int{
			{2, 2, 2},
			{2, 2, 0},
			{2, 0, 0},
		}
		So(floodFill(image, 1, 1, 2), ShouldResemble, res)
		//So(, ShouldEqual, -1)
	})
	Convey("2", t, func() {
		image := [][]int{
			{0, 0, 0},
			{0, 0, 0},
		}
		res := [][]int{
			{2, 2, 2},
			{2, 2, 2},
		}
		So(floodFill(image, 0, 0, 2), ShouldResemble, res)
		//So(, ShouldEqual, -1)
	})
}
