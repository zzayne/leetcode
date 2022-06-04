package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_merge(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			m1 := []int{1, 2, 3, 0, 0, 0}
			merge(m1, 3, []int{2, 5, 6}, 3)
			So(m1, ShouldResemble, []int{1, 2, 2, 3, 5, 6})
		})
		Convey("2", func() {
			m1 := []int{0}
			merge(m1, 0, []int{1}, 1)
			So(m1, ShouldResemble, []int{1})
		})
	})
}
