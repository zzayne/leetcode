package array

import (
	"testing"
)
import (
	. "github.com/smartystreets/goconvey/convey"
)

func Test_fourSum(t *testing.T) {
	Convey("test1", t, func() {
		So(fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0), ShouldResemble, [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}})
		So(fourSum([]int{1, 0, -1, 0, -2, 2}, 0), ShouldResemble, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}})
		So(fourSum([]int{2, 2, 2, 2, 2}, 8), ShouldResemble, [][]int{{2, 2, 2, 2}})
		///So(threeSum([]int{0, 1, 1}), ShouldResemble, nil)
		//So(, ShouldEqual, -1)
	})

}
