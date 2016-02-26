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
	"pkg.deepin.io/lib/graphic"
	"strings"

	"../theme/config"
)

type wallpaperPkgCreator struct {
	wpg *wallpaperGenerator
}

func (c *wallpaperPkgCreator) GetURL(id string) (string, error) {
	return "", fmt.Errorf("No Such URL")
}

func (c *wallpaperPkgCreator) Get(id string) (io.ReadCloser, error) {
	buf := new(bytes.Buffer)
	gw := gzip.NewWriter(buf)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	cfg, err := c.wpg.creators["config"].Get(id)
	if nil != err {
		return nil, err
	}
	defer cfg.Close()

	cfgTarName := "Wallpaper/.meta/theme.ini"
	cfgBody, err := ioutil.ReadAll(cfg)
	if nil != err {
		return nil, err
	}
	tw.WriteHeader(&tar.Header{
		Name: cfgTarName,
		Size: int64(len(cfgBody)),
		Mode: 0644,
	})
	tw.Write(cfgBody)

	wp, err := config.ReadWallpaperConfigString(string(cfgBody))
	if nil != err {
		return nil, err
	}

	for _, v := range strings.Split(wp.Extension.Ids, ";") {
		if 0 == len(v) {
			continue
		}
		d, err := c.wpg.creators["data"].Get(v)
		if nil != err {
			return nil, err
		}
		bkName := fmt.Sprint(wp.Get(v, "Name"))
		if "" == bkName {
			bkName = v
		}
		dataTarName := "Wallpaper/" + bkName
		dataBody, err := ioutil.ReadAll(d)
		if nil != err {
			return nil, err
		}
		tw.WriteHeader(&tar.Header{
			Name: dataTarName,
			Size: int64(len(dataBody)),
			Mode: 0644,
		})
		tw.Write(dataBody)
	}

	tw.Close()
	gw.Close()

	return ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func (c *wallpaperPkgCreator) Put(id string, r io.Reader) error {
	return upform.SlicePostData(r, (id))
}

type wallpaperGenerator struct {
	client   *http.Client
	Type     string
	creators map[string]creator
}

func newWallpaperGenrator() generator {
	g := &wallpaperGenerator{
		client:   &http.Client{},
		creators: map[string]creator{},
	}

	tpls := map[string]string{
		"meta":         "/subtheme/wallpaper/%s/meta.tar.gz",
		"config":       "/subtheme/wallpaper/%s/theme.ini",
		"data":         "/data/wallpaper/%s",
		"data-preview": "/data/wallpaper/%s-thumbnail",
	}

	for k, v := range tpls {
		g.creators[k] = &urlCreator{
			client:      g.client,
			urlTemplate: v,
		}
	}

	g.creators["package"] = &wallpaperPkgCreator{
		wpg: g,
	}
	return g
}

func (g *wallpaperGenerator) GetURL(datatype string, id string) (string, error) {
	return g.creators[datatype].GetURL(id)
}

func (wpg *wallpaperGenerator) Get(datatype string, id string) (io.ReadCloser, error) {
	return wpg.creators[datatype].Get(id)
}

func (wpg *wallpaperGenerator) Put(datatype string, r io.Reader) error {
	//parse reader
	gr, _ := gzip.NewReader(r)
	defer gr.Close()
	tr := tar.NewReader(gr)
	//Just Extract to tmpDir
	tmpPath := randTmpPath()
	if err := Extrat(tr, tmpPath); nil != err {
		return err
	}
	return wpg.putDir(tmpPath)
}

func (wpg *wallpaperGenerator) putDir(rootPath string) error {
	//Put the theme.ini
	configPath := rootPath + "/Wallpaper/.meta/theme.ini"
	configFile, err := os.Open(configPath)
	if nil != err {
		return err
	}
	defer configFile.Close()

	cfgBody, err := ioutil.ReadAll(configFile)
	if nil != err {
		return err
	}

	wp, err := config.ReadWallpaperConfigString(string(cfgBody))
	if nil != err {
		return err
	}

	//Put the config file
	if err := wpg.creators["config"].Put(wp.Theme.Id, bytes.NewReader(cfgBody)); nil != err {
		return err
	}

	for _, v := range strings.Split(wp.Extension.Ids, ";") {
		if 0 == len(v) {
			continue
		}
		wpdataPath := rootPath + "/Wallpaper/" + fmt.Sprint(wp.Get(v, "Name"))
		wpdata, err := os.Open(wpdataPath)
		if nil != err {
			return err
		}
		defer wpdata.Close()

		err = wpg.creators["data"].Put(v, wpdata)
		if nil != err {
			return err
		}
		//create thunbnial
		thumbnialPath := wpdataPath + "thumbnial"
		if err := graphic.ThumbnailImage(wpdataPath, thumbnialPath, 128, 72, graphic.FormatPng); nil != err {
			return err
		}
		thb, err := os.Open(thumbnialPath)
		if nil != err {
			return err
		}
		defer thb.Close()
		err = wpg.creators["data-preview"].Put(v, thb)
		if nil != err {
			return err
		}
	}

	//updata .meta file
	buf, err := Package("Wallpaper/.meta/", rootPath+"/Wallpaper/.meta/")
	if nil != err {
		return err
	}
	if err := wpg.creators["meta"].Put(wp.Theme.Id, bytes.NewReader(buf.Bytes())); nil != err {
		return err
	}

	return nil
}
