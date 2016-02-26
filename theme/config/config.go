package config

import (
	"github.com/Iceyer/gkv"
)

type themeInfo struct {
	Id          string
	Author      string
	Email       string
	Description string
	Name        string
}

type subthemes struct {
	Widget    string
	Icon      string
	Cursor    string
	Font      string
	Wallpaper string
}

type Theme struct {
	Theme     themeInfo
	SubThemes subthemes
}

func ReadThemeConfigFile(filepath string) (*Theme, error) {
	t := &Theme{}
	return t, gkv.ReadFileInto(t, filepath)
}

type subthemeInfo struct {
	Id              string
	Author          string
	Email           string
	Type            string
	Description     string
	Name            string
	SearchMethod    string
	SearchArguments string
}

type Widget struct {
	Theme subthemeInfo
}

type Icon struct {
	Theme subthemeInfo
}

type Cursor struct {
	Theme subthemeInfo
}

type SubthemeConfig struct {
	Theme subthemeInfo
}

func ReadSubthemeConfigFile(filepath string) (*SubthemeConfig, error) {
	sbt := &SubthemeConfig{}
	return sbt, gkv.ReadFileInto(sbt, filepath)
}

func ReadWidgetConfigFile(filepath string) (*Widget, error) {
	w := &Widget{}
	return w, gkv.ReadFileInto(w, filepath)
}

func ReadIconConfigFile(filepath string) (*Icon, error) {
	i := &Icon{}
	return i, gkv.ReadFileInto(i, filepath)
}

func ReadCursorConfigFile(filepath string) (*Cursor, error) {
	c := &Cursor{}
	return c, gkv.ReadFileInto(c, filepath)
}

type fontExt struct {
	Standard  string
	Monospace string
	Size      int
}

type Font struct {
	gkv.Config
	Theme     subthemeInfo
	Extension fontExt
}

type wallpaperExt struct {
	Default string
	Ids     string
}

type Wallpaper struct {
	gkv.Config
	Theme     subthemeInfo
	Extension wallpaperExt
}

func ReadFontConfigFile(filepath string) (*Font, error) {
	f := &Font{}
	return f, gkv.ReadFileInto(f, filepath)
}

func ReadFontConfigString(str string) (*Font, error) {
	f := &Font{}
	return f, gkv.ReadStringInto(f, str)
}

func ReadWallpaperConfigFile(filepath string) (*Wallpaper, error) {
	w := &Wallpaper{}
	return w, gkv.ReadFileInto(w, filepath)
}

func ReadWallpaperConfigString(str string) (*Wallpaper, error) {
	w := &Wallpaper{}
	return w, gkv.ReadStringInto(w, str)
}
