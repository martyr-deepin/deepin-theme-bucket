package bucket

import (
	"../theme/config"
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type wallpaperPkgCreator struct {
	wpg *wallpaperGenerator
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
		"config":       "/subtheme/wallpaper/%s/theme.ini",
		"data":         "/data/wallpaper/%s",
		"data-preview": "/data/wallpaper/%s-thumbnail-128x72",
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

func (wpg *wallpaperGenerator) Get(datatype string, id string) (io.ReadCloser, error) {
	return wpg.creators[datatype].Get(id)
}
