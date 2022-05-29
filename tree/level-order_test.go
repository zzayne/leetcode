package tree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			head := GetTree([]*int{P(3), P(9), P(20), nil, nil, P(15), P(17)})
			res := [][]int{
				{3},
				{9, 20},
				{15, 17},
			}
			So(levelOrder(head), ShouldResemble, res)
		})
		Convey("2", func() {
			head := GetTree([]*int{P(1)})
			res := [][]int{
				{1},
			}
			So(levelOrder(head), ShouldResemble, res)
		})
		Convey("3", func() {
			head := GetTree([]*int{nil})
			So(len(levelOrder(head)), ShouldEqual, 0)
		})
	})
}
