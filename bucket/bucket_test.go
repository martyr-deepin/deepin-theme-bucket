package bucket

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"os"
	"testing"
)

func startSever() {
	http.Handle("/", http.FileServer(http.Dir("./testdata/bucket")))
	http.ListenAndServe(":8300", nil)
}

func TestGet(t *testing.T) {
	go startSever()
	b := NewBucket()
	BucketHost = "http://127.0.0.1:8300/"
	tmpPath := os.TempDir() + "/theme-bucket"
	os.MkdirAll(tmpPath, 0755)
	tmpFile := tmpPath + "/.tmpFile"

	Convey("Test Theme Resource Get", t, func() {
		Convey("Test Theme Get", func() {
			err := b.GetFile("theme", "config", "32cb37c5e2ad6d937838f7f1e6431aef", tmpPath+"/theme-config.ini")
			So(err, ShouldBeNil)

			err = b.GetFile("theme", "package", "32cb37c5e2ad6d937838f7f1e6431aef", tmpPath+"/theme-package.tar.gz")
			So(err, ShouldBeNil)

			err = b.GetFile("theme", "thumbnail", "32cb37c5e2ad6d937838f7f1e6431aef", tmpPath+"/theme-thumbnail.png")
			So(err, ShouldBeNil)

			err = b.GetFile("theme", "icon-preview", "32cb37c5e2ad6d937838f7f1e6431aef", tmpPath+"/icon-preview.png")
			So(err, ShouldBeNil)
		})

		Convey("Test Widget Get", func() {
			err := b.GetFile("widget", "config", "25455af59da824bb40a6887ac1c38384", tmpFile)
			So(err, ShouldBeNil)

			err = b.GetFile("widget", "package", "25455af59da824bb40a6887ac1c38384", tmpFile)
			So(err, ShouldBeNil)

			err = b.GetFile("widget", "thumbnail", "25455af59da824bb40a6887ac1c38384", tmpFile)
			So(err, ShouldBeNil)
		})

		Convey("Test Icon Get", func() {
			err := b.GetFile("icon", "config", "4c29a66edd0ea62be95057ebed2d89c9", tmpFile)
			So(err, ShouldBeNil)

			err = b.GetFile("icon", "package", "4c29a66edd0ea62be95057ebed2d89c9", tmpFile)
			So(err, ShouldBeNil)

			err = b.GetFile("icon", "thumbnail", "4c29a66edd0ea62be95057ebed2d89c9", tmpFile)
			So(err, ShouldBeNil)
		})

		Convey("Test Cursor Get", func() {
			err := b.GetFile("cursor", "config", "19f78f43fb225a928912aed87051bd40", tmpFile)
			So(err, ShouldBeNil)

			err = b.GetFile("cursor", "package", "19f78f43fb225a928912aed87051bd40", tmpFile)
			So(err, ShouldBeNil)

			err = b.GetFile("cursor", "thumbnail", "19f78f43fb225a928912aed87051bd40", tmpFile)
			So(err, ShouldBeNil)
		})

		Convey("Test Wallpaper Get", func() {
			err := b.GetFile("wallpaper", "config", "bc7079cf32f031f4ceeee2beb10e5a07", tmpPath+"/config.ini")
			So(err, ShouldBeNil)

			err = b.GetFile("wallpaper", "data", "fffc4db6d2d5a94a9fd4d26ce959b744", tmpPath+"/a.png")
			So(err, ShouldBeNil)

			err = b.GetFile("wallpaper", "package", "bc7079cf32f031f4ceeee2beb10e5a07", tmpPath+"/wallpaper.tar.gz")
			So(err, ShouldBeNil)
		})

		Convey("Test Font Get", func() {
			err := b.GetFile("font", "config", "b97fd65d12519a1bba33077f76f32c2c", tmpPath+"/config.ini")
			So(err, ShouldBeNil)

			err = b.GetFile("font", "data", "43bb4cbf1d0ecfdb1309e4cb67264f35.ttf", tmpPath+"/SourceCodePro-Regular.ttf")
			So(err, ShouldBeNil)

			err = b.GetFile("font", "package", "b97fd65d12519a1bba33077f76f32c2c", tmpPath+"/font.tar.gz")
			So(err, ShouldBeNil)
		})
	})
}
