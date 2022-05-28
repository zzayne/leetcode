package tree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func P(i int) *int {
	return &i
}

func Test_maxDepth(t *testing.T) {
	//[3,9,20,null,null,15,7]，

	Convey("test2", t, func() {
		head := GetTree([]*int{P(3), P(9), P(20), nil, nil, P(15), P(7)})
		So(maxDepth(head), ShouldEqual, 3)
	})

	//Convey("test3", t, func() {
	//	So(firstUniqChar("aabb"), ShouldEqual, -1)
	//})
}
