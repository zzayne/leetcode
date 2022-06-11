package cqueue

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test(t *testing.T) {
	Convey("test1", t, func() {
		obj := Constructor()
		obj.AppendTail(3)
		val := obj.DeleteHead()
		So(val, ShouldEqual, 3)

		val = obj.DeleteHead()
		So(val, ShouldEqual, -1)
	})

	Convey("test2", t, func() {
		obj := Constructor()
		val := obj.DeleteHead()
		So(val, ShouldEqual, -1)

		obj.AppendTail(5)
		obj.AppendTail(2)

		val = obj.DeleteHead()
		So(val, ShouldEqual, 5)

		val = obj.DeleteHead()
		So(val, ShouldEqual, 2)
	})
}
