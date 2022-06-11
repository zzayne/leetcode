package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			m1 := []int{-4, -1, 0, 3, 10}
			So(sortedSquares(m1), ShouldResemble, []int{0, 1, 9, 16, 100})
		})
		Convey("2", func() {
			m1 := []int{-7, -3, 2, 3, 11}
			So(sortedSquares(m1), ShouldResemble, []int{4, 9, 9, 49, 121})
		})
	})
}
