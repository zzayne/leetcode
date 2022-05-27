package list

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			node := getCycleNodeList([]int{3, 2, 0, -4}, 1)
			So(hasCycle(node), ShouldEqual, true)
		})
		Convey("2", func() {
			node := getCycleNodeList([]int{3, 2, 0, -4}, 2)
			So(hasCycle(node), ShouldEqual, true)
		})
		Convey("3", func() {
			node := getCycleNodeList([]int{1, 2}, 0)
			So(hasCycle(node), ShouldEqual, true)
		})
		Convey("4", func() {
			node := getCycleNodeList([]int{1, 2}, 1)
			So(hasCycle(node), ShouldEqual, true)
		})
		Convey("5", func() {
			node := getCycleNodeList(nil, -1)
			So(hasCycle(node), ShouldEqual, false)
		})

	})
}
