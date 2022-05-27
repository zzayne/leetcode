package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_strStr(t *testing.T) {
	Convey("test1", t, func() {

		//Convey("1", func() {
		//	So(strStr("hello", "ll"), ShouldEqual, 2)
		//})
		//Convey("2", func() {
		//	So(strStr("aaaaa", "bba"), ShouldEqual, -1)
		//})
		//Convey("3", func() {
		//	So(strStr("aa", "bbabbb"), ShouldEqual, -1)
		//})
		//Convey("4", func() {
		//	So(strStr("a", "a"), ShouldEqual, 0)
		//})
		Convey("5", func() {
			So(strStr("abc", "c"), ShouldEqual, 2)
		})
	})
}
