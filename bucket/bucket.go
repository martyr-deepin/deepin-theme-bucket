package bucket

import (
	"io"
	"os"
)

type Bucket struct {
	gen map[string]generator
}

func NewBucket() *Bucket {
	return &Bucket{
		gen: map[string]generator{
			"wallpaper": newWallpaperGenrator(),
			"font":      newFontGenrator(),
			"theme":     newThemeGenrator(),
			"widget":    newWidgetGenrator(),
			"icon":      newIconGenrator(),
			"cursor":    newCursorGenrator(),
		},
	}
}

func (b *Bucket) Get(themeType string, datatype string, id string) (io.ReadCloser, error) {
	return b.gen[themeType].Get(datatype, id)
}

func (b *Bucket) GetFile(themeType string, datatype string, id string, filepath string) error {
	rd, err := b.gen[themeType].Get(datatype, id)
	if nil != err {
		return err
	}
	defer rd.Close()

	f, err := os.Create(filepath)
	if nil != err {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, rd)
	if nil != err {
		return err
	}
	return nil
}
