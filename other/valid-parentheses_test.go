package other

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isValid(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(isValid("()"), ShouldEqual, true)
		})
		Convey("2", func() {
			So(isValid("()[]{}"), ShouldEqual, true)
		})
		Convey("3", func() {
			So(isValid("(]"), ShouldEqual, false)
		})
		Convey("4", func() {
			So(isValid("([)]"), ShouldEqual, false)
		})
		Convey("5", func() {
			So(isValid("{[]}"), ShouldEqual, true)
		})
		Convey("6", func() {
			So(isValid("{[[(({}))]]}"), ShouldEqual, true)
		})
		Convey("7", func() {
			So(isValid("()))"), ShouldEqual, false)
		})
	})
}
