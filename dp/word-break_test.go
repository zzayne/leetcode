package dp

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(wordBreak("leetcode", []string{"leet", "code"}), ShouldEqual, true)
		})
		Convey("2", func() {
			So(wordBreak("applepenapple", []string{"apple", "pen"}), ShouldEqual, true)
		})
		Convey("3", func() {
			So(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), ShouldEqual, false)
		})
		Convey("4", func() {
			So(wordBreak("aaaaaaa", []string{"aaaa", "aa"}), ShouldEqual, false)
		})
	})
}
