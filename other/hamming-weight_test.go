package other

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(hammingWeight(00000000000000000000000000001011), ShouldEqual, 3)
		})
		Convey("2", func() {
			So(hammingWeight(00000000000000000000000010000000), ShouldEqual, 1)
		})
	})
}
