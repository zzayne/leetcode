package tree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isSymmetric(t *testing.T) {
	Convey("test1", t, func() {
		head := GetTree([]*int{P(1), P(2), P(2), P(3), P(4), P(4), P(3)})
		So(isSymmetric(head), ShouldEqual, true)
	})
}
