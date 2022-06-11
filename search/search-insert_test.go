package search

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(searchInsert([]int{1, 3, 5, 6}, 5), ShouldEqual, 2)
		})
		Convey("2", func() {
			So(searchInsert([]int{1, 3, 5, 6}, 2), ShouldEqual, 1)
		})
		Convey("3", func() {
			So(searchInsert([]int{1, 3, 5, 6}, 7), ShouldEqual, 4)
		})

	})
}
