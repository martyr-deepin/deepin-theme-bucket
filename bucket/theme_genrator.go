package bucket

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"../theme/config"
)

type themeGenerator struct {
	client       *http.Client
	creators     map[string]creator
	subthemeType string
}

func (g *themeGenerator) GetURL(datatype string, id string) (string, error) {
	return g.creators[datatype].GetURL(id)
}

func (g *themeGenerator) Get(datatype string, id string) (io.ReadCloser, error) {
	return g.creators[datatype].Get(id)
}

func newThemeGenrator() generator {
	g := &themeGenerator{
		client:   &http.Client{},
		creators: map[string]creator{},
	}

	tpls := map[string]string{
		"meta":           "/theme/%s/meta.tar.gz",
		"config":         "/theme/%s/theme.ini",
		"thumbnail":      "/theme/%s/thumbnail.png",
		"package":        "/theme/%s/package.tar.gz",
		"cursor-preview": "/theme/%s/preview/cursor.png",
		"icon-preview":   "/theme/%s/preview/icon.png",
		"widget-preview": "/theme/%s/preview/widget.png",
	}

	for k, v := range tpls {
		g.creators[k] = &urlCreator{
			client:      g.client,
			urlTemplate: v,
		}
	}

	return g
}

func (tg *themeGenerator) Put(datatype string, r io.Reader) error {
	//parse reader
	gr, _ := gzip.NewReader(r)
	defer gr.Close()
	tr := tar.NewReader(gr)
	//Just Extract to tmpDir
	tmpPath := randTmpPath()
	if err := Extrat(tr, tmpPath); nil != err {
		return err
	}
	return tg.putDir(tmpPath)
}

func (tg *themeGenerator) putDir(rootPath string) error {
	files, _ := ioutil.ReadDir(rootPath)
	if 1 != len(files) {
		return fmt.Errorf("Package Struct Error, Must One Dir at root.")
	}
	themeName := files[0].Name()
	themeRoot := rootPath + "/" + themeName

	//put config file
	cfgPath := themeRoot + "/theme.ini"
	cfg, err := config.ReadThemeConfigFile(cfgPath)
	if nil != err {
		return err
	}

	filelist := map[string]string{
		"config":         cfgPath,
		"thumbnail":      themeRoot + "/thumbnail.png",
		"cursor-preview": themeRoot + "/Preview/cursor.png",
		"icon-preview":   themeRoot + "/Preview/icon.png",
		"widget-preview": themeRoot + "/Preview/widget.png",
	}
	if err := putFileList(cfg.Theme.Id, filelist, &tg.creators); nil != err {
		return err
	}

	//put subtheme
	genlist := []generator{newWallpaperGenrator(), newFontGenrator(), newCursorGenrator(), newWidgetGenrator(), newIconGenrator()}
	for _, g := range genlist {
		if err := g.putDir(themeRoot); nil != err {
			return err
		}
	}

	//put metafile
	sublist := []string{"Wallpaper", "Font", "Icon", "Widget", "Cursor"}
	for _, s := range sublist {
		os.RemoveAll(themeRoot + "/" + s)
	}

	buf, err := Package("", rootPath)
	if nil != err {
		return err
	}
	if err := tg.creators["meta"].Put(cfg.Theme.Id, bytes.NewReader(buf.Bytes())); nil != err {
		return err
	}
	return nil
}

func putFileList(id string, filelist map[string]string, creators *map[string]creator) error {
	for k, v := range filelist {
		f, err := os.Open(v)
		if nil != err {
			return err
		}
		defer f.Close()
		if err := (*creators)[k].Put(id, f); nil != err {
			return err
		}
	}
	return nil
}
