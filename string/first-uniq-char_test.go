package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_firstUniqChar(t *testing.T) {
	//Convey("test1", t, func() {
	//	So(firstUniqChar("leetcode"), ShouldEqual, 0)
	//})

	Convey("test2", t, func() {
		So(firstUniqChar("aadadaad"), ShouldEqual, -1)
	})

	//Convey("test3", t, func() {
	//	So(firstUniqChar("aabb"), ShouldEqual, -1)
	//})
}
