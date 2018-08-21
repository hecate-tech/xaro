package player

import (
	"errors"
	"log"
	"testing"

	"engo.io/ecs"
	"engo.io/engo/common"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	defer func() {
		loadedSprite = common.LoadedSprite
		logFatalln = log.Fatalln
	}()
	w := &ecs.World{}

	Convey("Given I call New", t, func() {
		loadedSprite = func(string) (*common.Texture, error) { return &common.Texture{}, nil }
		player := New(w)
		Convey("Then a new player object should not be nil", func() {
			So(player, ShouldNotBeNil)
		})
	})
	Convey("Given I call New with a bad Sprite", t, func() {
		var err error
		logFatalln = func(...interface{}) { err = errors.New("PASSED") }
		loadedSprite = func(string) (*common.Texture, error) { return &common.Texture{}, errors.New("FAIL") }
		New(w)
		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
}

func TestUpdate(t *testing.T) {

}

func TestRemove(t *testing.T) {

}
