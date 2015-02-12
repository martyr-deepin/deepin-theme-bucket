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

func (b *Bucket) put(themeType string, r io.Reader) error {
	return b.gen[themeType].Put("package", r)
}

func (b *Bucket) putFile(themeType string, filepath string) error {
	return nil
}

func (b *Bucket) getURL(themeType string, datatype string, id string) (string, error) {
	return b.gen[themeType].GetURL(datatype, id)
}

func (b *Bucket) getFile(themeType string, datatype string, id string) (string, error) {
	//create new path
	filepath := randTmpPath()
	f, err := os.Create(filepath)
	if nil != err {
		return "", err
	}
	defer f.Close()

	rd, err := b.gen[themeType].Get(datatype, id)
	if nil != err {
		return "", err
	}
	defer rd.Close()

	_, err = io.Copy(f, rd)
	if nil != err {
		return "", err
	}
	return filepath, nil
}
