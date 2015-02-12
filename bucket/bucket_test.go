package bucket

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"pkg.linuxdeepin.com/lib/utils"
	"testing"
)

var (
	b                      = NewBucket()
	LoaclBucketHost        = "http://127.0.0.1:8300"
	theme_meta_template    = LoaclBucketHost + "/theme/%s/meta.tar.gz"
	subtheme_meta_template = LoaclBucketHost + "/subtheme/%s/%s/meta.tar.gz"
	theme_pkg_template     = LoaclBucketHost + "/theme/%s/package.tar.gz"
	subtheme_pkg_template  = LoaclBucketHost + "/subtheme/%s/%s/package.tar.gz"
	data_template          = LoaclBucketHost + "/data/%s/%s"
)

func startSever() {
	http.Handle("/", http.FileServer(http.Dir("./testdata/bucket")))
	http.ListenAndServe(":8300", nil)
}

func TestGetFile(t *testing.T) {
	BucketHost = LoaclBucketHost
	go startSever()

	Convey("Test Get Theme Resource File", t, func() {
		Convey("Test Theme File Get", func() {
			themetype := "theme"
			id := "32cb37c5e2ad6d937838f7f1e6431aef"

			url, err := b.GetMetaFile(themetype, id)
			So(err, ShouldBeNil)
			md5, _ := utils.SumFileMd5(url)
			So(md5, ShouldEqual, "7af6705b5fa6cd631ec175e397d4e41e")

			url, err = b.GetThemePackageFile(id)
			So(err, ShouldBeNil)
			md5, _ = utils.SumFileMd5(url)
			So(md5, ShouldEqual, "7af6705b5fa6cd631ec175e397d4e41e")
		})

		Convey("Test Widget Get", func() {
			themetype := "widget"
			id := "25455af59da824bb40a6887ac1c38384"

			url, err := b.GetMetaFile(themetype, id)
			So(err, ShouldBeNil)
			md5, _ := utils.SumFileMd5(url)
			So(md5, ShouldEqual, "9594ba4c1be09d44ac2b9d54d42315f3")

			url, err = b.GetWidgetPackageFile(id)
			So(err, ShouldBeNil)
			md5, _ = utils.SumFileMd5(url)
			So(md5, ShouldEqual, "968853d5fd2d971f7d83c6c0018e030a")
		})

		Convey("Test Icon Get", func() {
			themetype := "icon"
			id := "4c29a66edd0ea62be95057ebed2d89c9"

			url, err := b.GetMetaFile(themetype, id)
			So(err, ShouldBeNil)
			md5, _ := utils.SumFileMd5(url)
			So(md5, ShouldEqual, "09fc1b2e0b89f8e10d5e5d38e8ecb34d")

			url, err = b.GetIconPackageFile(id)
			So(err, ShouldBeNil)
			md5, _ = utils.SumFileMd5(url)
			So(md5, ShouldEqual, "f2e0ccc7071284ec96f739af0c8fe7f2")
		})

		Convey("Test Cursor Get", func() {
			themetype := "cursor"
			id := "19f78f43fb225a928912aed87051bd40"

			url, err := b.GetMetaFile(themetype, id)
			So(err, ShouldBeNil)
			md5, _ := utils.SumFileMd5(url)
			So(md5, ShouldEqual, "af53b817be3ee112ee8ab224dda1d1e3")

			url, err = b.GetCursorPackageFile(id)
			So(err, ShouldBeNil)
			md5, _ = utils.SumFileMd5(url)
			So(md5, ShouldEqual, "837ef42a5405a91bdea7452a33f55f38")
		})

		Convey("Test Wallpaper Get", func() {
			themetype := "wallpaper"
			id := "bc7079cf32f031f4ceeee2beb10e5a07"

			url, err := b.GetMetaFile(themetype, id)
			So(err, ShouldBeNil)
			md5, _ := utils.SumFileMd5(url)
			So(md5, ShouldEqual, "a3ce1fbbc3e5ebbc9873ab9a308d2075")

			url, err = b.GetWallpaperPackageFile(id)
			So(err, ShouldBeNil)
			md5, _ = utils.SumFileMd5(url)
			So(md5, ShouldEqual, "545d506784130f7c28b4058a47d75afe")
		})

		Convey("Test Font Get", func() {
			themetype := "font"
			id := "b97fd65d12519a1bba33077f76f32c2c"

			url, err := b.GetMetaFile(themetype, id)
			So(err, ShouldBeNil)
			md5, _ := utils.SumFileMd5(url)
			So(md5, ShouldEqual, "01ba7ea2d21b303f2a72bab643967147")

			url, err = b.GetFontPackageFile(id)
			So(err, ShouldBeNil)
			md5, _ = utils.SumFileMd5(url)
			So(md5, ShouldEqual, "0701b3ed1265becd993a2f1f7d40b3e8")
		})
	})
}

func TestGetURL(t *testing.T) {
	BucketHost = LoaclBucketHost
	Convey("Test Theme Resource Get", t, func() {
		Convey("Test Theme Get", func() {
			themetype := "theme"
			id := "32cb37c5e2ad6d937838f7f1e6431aef"

			url, err := b.GetMeta(themetype, id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(theme_meta_template, id))

			url, err = b.GetThemePackage(id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(theme_pkg_template, id))
		})

		Convey("Test Widget Get", func() {
			themetype := "widget"
			id := "25455af59da824bb40a6887ac1c38384"

			url, err := b.GetMeta(themetype, id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(subtheme_meta_template, themetype, id))

			url, err = b.GetWidgetPackage(id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(subtheme_pkg_template, themetype, id))
		})

		Convey("Test Icon Get", func() {
			themetype := "icon"
			id := "4c29a66edd0ea62be95057ebed2d89c9"

			url, err := b.GetMeta(themetype, id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(subtheme_meta_template, themetype, id))

			url, err = b.GetIconPackage(id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(subtheme_pkg_template, themetype, id))
		})

		Convey("Test Cursor Get", func() {
			themetype := "cursor"
			id := "19f78f43fb225a928912aed87051bd40"

			url, err := b.GetMeta(themetype, id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(subtheme_meta_template, themetype, id))

			url, err = b.GetCursorPackage(id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(subtheme_pkg_template, themetype, id))
		})

		Convey("Test Wallpaper Get", func() {
			themetype := "wallpaper"
			id := "bc7079cf32f031f4ceeee2beb10e5a07"

			url, err := b.GetMeta(themetype, id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(subtheme_meta_template, themetype, id))

			url, err = b.GetWallpaper(id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(data_template, themetype, id))
		})

		Convey("Test Font Get", func() {
			themetype := "font"
			id := "b97fd65d12519a1bba33077f76f32c2c"

			url, err := b.GetMeta(themetype, id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(subtheme_meta_template, themetype, id))

			url, err = b.GetFont(id)
			So(err, ShouldBeNil)
			So(url, ShouldEqual, fmt.Sprintf(data_template, themetype, id))
		})
	})
}
