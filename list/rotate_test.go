package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_rotate(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			m1 := []int{1, 2, 3, 4, 5, 6, 7}
			rotate(m1, 3)
			So(m1, ShouldResemble, []int{5, 6, 7, 1, 2, 3, 4})
		})
		Convey("2", func() {
			m1 := []int{-1, -100, 3, 99}
			rotate(m1, 2)
			So(m1, ShouldResemble, []int{3, 99, -1, -100})
		})
	})
}
