package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	Convey("Test_isAnagram", t, func() {
		Convey("1", func() {
			So(isAnagram("anagram", "nagaram"), ShouldEqual, true)
		})
		Convey("2", func() {
			So(isAnagram("anagram", "aaanagaram"), ShouldEqual, false)
		})
		Convey("3", func() {
			So(isAnagram("rat", "car"), ShouldEqual, false)
		})
		Convey("4", func() {
			So(isAnagram("aaaa", "aaaa"), ShouldEqual, true)
		})
	})

}
