package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_reverseString(t *testing.T) {
	Convey("test1", t, func() {
		in := []byte{'h', 'e', 'l', 'l', 'o'}
		reverseString(in)

		So(in, ShouldResemble, []byte{'o', 'l', 'l', 'e', 'h'})
	})
}
