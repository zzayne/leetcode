package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_threeSum(t *testing.T) {
	Convey("test1", t, func() {

		So(threeSum([]int{-1, 0, 1, 2, -1, -4}), ShouldResemble, [][]int{{-1, -1, 2}, {-1, 0, 1}})
		So(threeSum([]int{0, 0, 0, 0}), ShouldResemble, [][]int{{0, 0, 0}})
		So(threeSum([]int{1, -1, -1, 0}), ShouldResemble, [][]int{{-1, 0, 1}})
		//So(threeSum([]int{0, 1, 1}), ShouldResemble, nil)
		//So(, ShouldEqual, -1)
	})

}
