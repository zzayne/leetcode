package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_countAndSay(t *testing.T) {
	Convey("test1", t, func() {
		So(countAndSay(4), ShouldEqual, "1211")
		So(countAndSay(3), ShouldEqual, "21")
		//So(, ShouldEqual, -1)
	})
}
