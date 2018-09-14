package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewIDManager(t *testing.T) {
	Convey("Given NewIDManager is stored\n"+
		"into a variable", t, func() {

		idManager := NewIDManager()

		Convey("Then idManager should not be nil", func() {
			So(idManager, ShouldNotBeNil)
			So(idManager.playerIDs, ShouldNotBeNil)
			So(idManager.botsIDs, ShouldNotBeNil)
		})
	})
}

func TestNextPlayerID(t *testing.T) {
	Convey("Given a new IDManager", t, func() {
		idManager := NewIDManager()

		Convey("When I call NextPlayerID", func() {
			var id uint32
			Convey("Then the first time should equal 1", func() {
				id = idManager.NextPlayerID()
				So(id, ShouldEqual, 1)
			})

			idManager.NextPlayerID()
			Convey("Then the next time should equal 2", func() {
				id = idManager.NextPlayerID()
				So(id, ShouldEqual, 2)
			})

			idManager.NextPlayerID()
			Convey("Then the next time should equal 3", func() {
				id = idManager.NextPlayerID()
				So(id, ShouldEqual, 3)
			})
		})
	})
}
