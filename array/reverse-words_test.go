package array

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reverseWords(t *testing.T) {
	Convey("test1", t, func() {
		So(reverseWords("Let's take LeetCode contest"), ShouldEqual, "s'teL ekat edoCteeL tsetnoc")
		So(reverseWords("God Ding"), ShouldEqual, "doG gniD")
		//So(, ShouldEqual, -1)
	})
}
