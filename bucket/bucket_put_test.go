package bucket

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPutFile(t *testing.T) {
	b := NewBucket()

	Convey("Test Put Theme Resource File", t, func() {
		Convey("Test Theme File Put", func() {
			themetype := "theme"
			err := b.putFile(themetype, "testdata/Deepin.tar.gz")
			So(err, ShouldBeNil)
		})

		Convey("Test Wallpaper Get", func() {
			themetype := "wallpaper"
			err := b.putFile(themetype, "testdata/Wallpaper.tar.gz")
			So(err, ShouldBeNil)
		})

		Convey("Test Font Get", func() {
			themetype := "font"
			err := b.putFile(themetype, "testdata/Font.tar.gz")
			So(err, ShouldBeNil)

		})
	})
}
