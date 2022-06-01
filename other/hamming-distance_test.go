package other

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_hammingDistance(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(hammingDistance(1, 4), ShouldEqual, 2)
		})
		Convey("2", func() {
			So(hammingDistance(3, 1), ShouldEqual, 1)
		})
	})
}
