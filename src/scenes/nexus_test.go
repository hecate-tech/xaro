package scenes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPreload(t *testing.T) {
	n := Nexus{}

	Convey("Given we Preload Nexus", t, func() {
		Convey("Then Preload shouldn't panic", func() {
			So(n.Preload, ShouldNotPanic)
		})
	})
}
