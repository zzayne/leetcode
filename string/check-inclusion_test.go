package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	Convey("checkInclusion", t, func() {
		Convey("1", func() {
			So(checkInclusion("ab", "eidbaooo"), ShouldEqual, true)
		})
		Convey("2", func() {
			So(checkInclusion("ab", "eidboaoo"), ShouldEqual, false)
		})
		Convey("3", func() {
			So(checkInclusion("aba", "aaaabaabbb"), ShouldEqual, true)
		})
		Convey("4", func() {
			So(checkInclusion("adc", "dcda"), ShouldEqual, true)
		})
		Convey("5", func() {
			So(checkInclusion("hello", "ooolleoooleh"), ShouldEqual, false)
		})
		Convey("6", func() {
			So(checkInclusion("a", "ab"), ShouldEqual, true)
		})
	})
}
