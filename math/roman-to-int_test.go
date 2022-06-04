package math

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(romanToInt("III"), ShouldResemble, 3)
		})
		Convey("2", func() {
			So(romanToInt("IV"), ShouldResemble, 4)
		})
		Convey("3", func() {
			So(romanToInt("IX"), ShouldResemble, 9)
		})
		Convey("4", func() {
			So(romanToInt("LVIII"), ShouldResemble, 58)
		})
		Convey("5", func() {
			So(romanToInt("MCMXCIV"), ShouldResemble, 1994)
		})
	})
}
