package search

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_search(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(search([]int{-1, 0, 3, 5, 9, 12}, 9), ShouldEqual, 4)
		})
		Convey("2", func() {
			So(search([]int{-1, 0, 3, 5, 9, 12}, 2), ShouldEqual, -1)
		})

	})
}
