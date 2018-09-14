package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUserNameRegistry(t *testing.T) {
	Convey("Given NewUserNameRegistry is called", t, func() {
		userReg := NewUserNameRegistry()
		Convey("Then userReg should not be nil", func() {
			So(userReg, ShouldNotBeNil)
			So(userReg.userNames, ShouldNotBeNil)
		})
	})
}

func TestAddUserName(t *testing.T) {
	Convey("Given a username registry", t, func() {
		userReg := NewUserNameRegistry()
		Convey("When AddUserName is used", func() {
			userName := "Test"
			userReg.AddUserName(1, userName)
			Convey("Then userName should be added", func() {
				So(userReg.GetUserName(1), ShouldEqual, userName)
			})
		})
	})
}
