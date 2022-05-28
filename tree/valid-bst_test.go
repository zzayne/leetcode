package tree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	Convey("test1", t, func() {
		head := GetTree([]*int{P(2), P(1), P(5)})
		So(isValidBST(head), ShouldEqual, true)
	})
	Convey("test2", t, func() {
		head := GetTree([]*int{P(5), P(1), P(5), nil, nil, P(3), P(6)})
		So(isValidBST(head), ShouldEqual, false)
	})
	Convey("test3", t, func() {
		head := GetTree([]*int{P(0)})
		So(isValidBST(head), ShouldEqual, true)
	})
}
