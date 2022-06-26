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
		So(firstUniqChar("abaccdeff"), ShouldEqual, 'b')
	})
	Convey("test3", t, func() {
		So(firstUniqChar(""), ShouldEqual, ' ')
	})

	//Convey("test3", t, func() {
	//	So(firstUniqChar("aabb"), ShouldEqual, -1)
	//})
}
