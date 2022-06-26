package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_replaceSpace(t *testing.T) {
	Convey("test1", t, func() {

		So(replaceSpace("We are happy."), ShouldEqual, "We%20are%20happy.")
	})
}
