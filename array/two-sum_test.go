package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_twoSum(t *testing.T) {
	Convey("test1", t, func() {
		So(twoSum([]int{2, 7, 11, 15}, 9), ShouldResemble, []int{1, 2})
		So(twoSum([]int{2, 3, 4}, 6), ShouldResemble, []int{1, 3})
		//So(, ShouldEqual, -1)
	})
}
