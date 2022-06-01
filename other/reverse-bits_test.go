package other

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reverseBits(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(reverseBits(43261596), ShouldEqual, 964176192)
		})
		Convey("2", func() {
			So(reverseBits(4294967293), ShouldEqual, 3221225471)
		})
	})
}
