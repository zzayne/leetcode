package other

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_generateTriangle(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(generateTriangle(1), ShouldResemble, [][]int{{1}})
		})
		Convey("2", func() {
			So(generateTriangle(2), ShouldResemble, [][]int{{1}, {1, 1}})
		})
		Convey("3", func() {
			So(generateTriangle(3), ShouldResemble, [][]int{{1}, {1, 1}, {1, 2, 1}})
		})
		Convey("4", func() {
			So(generateTriangle(5), ShouldResemble, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}})
		})
	})
}
