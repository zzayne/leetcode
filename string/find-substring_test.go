package string

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	Convey("test1", t, func() {
		So(findSubstring("barfoothefoobarman", []string{"foo", "bar"}), ShouldResemble, []int{0, 9})
	})
}
