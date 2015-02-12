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
)

type fontPkgCreator struct {
	g *fontGenerator
}

func (c *fontPkgCreator) GetURL(id string) (string, error) {
	return "", fmt.Errorf("No Such URL")
}

func (c *fontPkgCreator) Get(id string) (io.ReadCloser, error) {
	buf := new(bytes.Buffer)
	gw := gzip.NewWriter(buf)
	defer gw.Close()

	tw := tar.NewWriter(buf)
	defer tw.Close()

	cfg, err := c.g.creators["config"].Get(id)
	if nil != err {
		return nil, err
	}
	defer cfg.Close()

	cfgTarName := "Font/.meta/theme.ini"
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

	fontCfg, err := config.ReadFontConfigString(string(cfgBody))
	if nil != err {
		return nil, err
	}

	fontList := []string{fontCfg.Extension.Standard, fontCfg.Extension.Monospace}
	//	fontCfg.Extension.Monospace
	for _, v := range fontList {
		fontFile := fmt.Sprint(fontCfg.Get(v, "File"))
		d, err := (c.g.creators["data"].Get(fontFile))
		if nil != err {
			return nil, err
		}
		defer d.Close()
		dataTarName := "Font/" + fontFile
		dataBody, err := ioutil.ReadAll(d)
		if nil != err {
			return nil, err
		}
		err = tw.WriteHeader(&tar.Header{
			Name: dataTarName,
			Size: int64(len(dataBody)),
			Mode: 0644,
		})
		if nil != err {
			return nil, err
		}
		_, err = tw.Write(dataBody)
		if nil != err {
			return nil, err
		}
	}

	tw.Close()
	gw.Close()
	return ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

type fontGenerator struct {
	client   *http.Client
	creators map[string]creator
}

func newFontGenrator() generator {
	g := &fontGenerator{
		client:   &http.Client{},
		creators: map[string]creator{},
	}

	tpls := map[string]string{
		"meta":   "/subtheme/font/%s/meta.tar.gz",
		"config": "/subtheme/font/%s/theme.ini",
		"data":   "/data/font/%s",
	}

	for k, v := range tpls {
		g.creators[k] = &urlCreator{
			client:      g.client,
			urlTemplate: v,
		}
	}

	g.creators["package"] = &fontPkgCreator{
		g: g,
	}
	return g
}

func (g *fontGenerator) GetURL(datatype string, id string) (string, error) {
	return g.creators[datatype].GetURL(id)
}

func (g *fontGenerator) Get(datatype string, id string) (io.ReadCloser, error) {
	return g.creators[datatype].Get(id)
}

func (wpg *fontGenerator) Put(datatype string, r io.Reader) error {
	return nil
}
