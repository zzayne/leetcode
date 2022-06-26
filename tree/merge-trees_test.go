package tree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	Convey("test1", t, func() {
		head := GetTree([]*int{P(3), P(9), P(20), nil, nil, P(15), P(7)})
		So(maxDepth(head), ShouldEqual, 3)
	})
}
