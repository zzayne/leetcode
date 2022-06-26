package math

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_fib(t *testing.T) {
	Convey("test1", t, func() {
		So(fib(3), ShouldResemble, 2)

	})
}
