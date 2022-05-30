package math

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(countPrimes(10), ShouldEqual, 4)
		})
		Convey("2", func() {
			So(countPrimes(0), ShouldEqual, 0)
		})
		Convey("3", func() {
			So(countPrimes(1), ShouldEqual, 0)
		})
	})
}
