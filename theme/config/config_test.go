package config

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGet(t *testing.T) {
	Convey("Test Theme Config File Read", t, func() {
		Convey("Test Theme Config", func() {
			theme, err := GetThemeConfig("testdata/theme.ini")
			So(err, ShouldBeNil)
			So(theme.Theme.Id, ShouldEqual, "32cb37c5e2ad6d937838f7f1e6431aef")
			So(theme.SubThemes.Widget, ShouldEqual, "25455af59da824bb40a6887ac1c38384")
			So(theme.SubThemes.Icon, ShouldEqual, "4c29a66edd0ea62be95057ebed2d89c9")
			So(theme.SubThemes.Cursor, ShouldEqual, "19f78f43fb225a928912aed87051bd40")
			So(theme.SubThemes.Font, ShouldEqual, "b97fd65d12519a1bba33077f76f32c2c")
			So(theme.SubThemes.Wallpaper, ShouldEqual, "bc7079cf32f031f4ceeee2beb10e5a07")
		})

		Convey("Test Icon Config", func() {
			icon, err := GetIconConfig("testdata/icon.ini")
			So(err, ShouldBeNil)
			So(icon.Theme.Id, ShouldEqual, "4c29a66edd0ea62be95057ebed2d89c9")
			So(icon.Theme.Name, ShouldEqual, "Deepin")
		})

		Convey("Test Font Config", func() {
			font, err := ReadFontConfigFile("testdata/font.ini")
			So(err, ShouldBeNil)
			So(font.Theme.Id, ShouldEqual, "b97fd65d12519a1bba33077f76f32c2c")
			So(font.Extension.Standard, ShouldEqual, "Source San Hans SC")
			So(font.Extension.Monospace, ShouldEqual, "Source Code Pro")
		})

		Convey("Test Wallpaper Config", func() {
			wp, err := ReadWallpaperConfigFile("testdata/wallpaper.ini")
			So(err, ShouldBeNil)
			So(wp.Theme.Id, ShouldEqual, "bc7079cf32f031f4ceeee2beb10e5a07")
			So(wp.Extension.Default, ShouldEqual, "de56a4e0fed1e29d0b6884a5305f92ff")
			So(wp.Extension.Ids, ShouldEqual, "fffc4db6d2d5a94a9fd4d26ce959b744;de56a4e0fed1e29d0b6884a5305f92ff;")
			So(fmt.Sprint(wp.Get(wp.Extension.Default, "Id")), ShouldEqual, "de56a4e0fed1e29d0b6884a5305f92ff")
		})
	})
}
