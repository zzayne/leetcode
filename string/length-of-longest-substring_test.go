package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	Convey("lengthOfLongestSubstring", t, func() {
		Convey("1", func() {
			So(lengthOfLongestSubstring("abcabcbb"), ShouldEqual, 3)
		})
		Convey("2", func() {
			So(lengthOfLongestSubstring("bbbbb"), ShouldEqual, 1)
		})
		Convey("3", func() {
			So(lengthOfLongestSubstring("pwwkew"), ShouldEqual, 3)
		})
	})
}
