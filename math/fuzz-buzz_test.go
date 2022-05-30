package math

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	Convey("test1", t, func() {
		Convey("1", func() {
			So(fuzzBuzz(3), ShouldResemble, []string{"1", "2", "Fizz"})
		})
		Convey("2", func() {
			So(fuzzBuzz(5), ShouldResemble, []string{"1", "2", "Fizz", "4", "Buzz"})
		})
	})
}
