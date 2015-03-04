package bucket

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"os"

	"../theme/config"
)

type subthemeGenerator struct {
	client       *http.Client
	creators     map[string]creator
	subthemeType string
}

func (g *subthemeGenerator) GetURL(datatype string, id string) (string, error) {
	return g.creators[datatype].GetURL(id)
}

func (g *subthemeGenerator) Get(datatype string, id string) (io.ReadCloser, error) {
	return g.creators[datatype].Get(id)
}

func newSubthemeGenrator(subthemeType string, urltpls map[string]string) generator {
	g := &subthemeGenerator{
		subthemeType: subthemeType,
		client:       &http.Client{},
		creators:     map[string]creator{},
	}
	for k, v := range urltpls {
		g.creators[k] = &urlCreator{
			client:      g.client,
			urlTemplate: v,
		}
	}

	return g
}

func newIconGenrator() generator {
	tpls := map[string]string{
		"meta":      "/subtheme/icon/%s/meta.tar.gz",
		"config":    "/subtheme/icon/%s/theme.ini",
		"thumbnail": "/subtheme/icon/%s/thumbnail.png",
		"package":   "/subtheme/icon/%s/package.tar.gz",
	}

	return newSubthemeGenrator("Icon", tpls)
}

func newWidgetGenrator() generator {
	tpls := map[string]string{
		"meta":      "/subtheme/widget/%s/meta.tar.gz",
		"config":    "/subtheme/widget/%s/theme.ini",
		"thumbnail": "/subtheme/widget/%s/thumbnail.png",
		"package":   "/subtheme/widget/%s/package.tar.gz",
	}
	return newSubthemeGenrator("Widget", tpls)
}

func newCursorGenrator() generator {
	tpls := map[string]string{
		"meta":      "/subtheme/cursor/%s/meta.tar.gz",
		"config":    "/subtheme/cursor/%s/theme.ini",
		"thumbnail": "/subtheme/cursor/%s/thumbnail.png",
		"package":   "/subtheme/cursor/%s/package.tar.gz",
	}
	return newSubthemeGenrator("Cursor", tpls)
}

func (stg *subthemeGenerator) Put(datatype string, r io.Reader) error {
	//parse reader
	gr, _ := gzip.NewReader(r)
	defer gr.Close()
	tr := tar.NewReader(gr)
	//Just Extract to tmpDir
	tmpPath := randTmpPath()
	if err := Extrat(tr, tmpPath); nil != err {
		return err
	}
	return stg.putDir(tmpPath)
}

func (stg *subthemeGenerator) putDir(rootPath string) error {
	//Put the theme.ini
	subthemeRoot := rootPath + "/" + stg.subthemeType + "/"
	configPath := subthemeRoot + ".meta/theme.ini"

	cfg, err := config.ReadSubthemeConfigFile(configPath)
	if nil != err {
		return err
	}

	configFile, err := os.Open(configPath)
	if nil != err {
		return err
	}
	defer configFile.Close()

	//Put config file
	if err := stg.creators["config"].Put(cfg.Theme.Id, configFile); nil != err {
		return err
	}

	//updata .meta file
	buf, err := Package(stg.subthemeType+"/.meta/", subthemeRoot+".meta/")
	if nil != err {
		return err
	}
	if err := stg.creators["meta"].Put(cfg.Theme.Id, bytes.NewReader(buf.Bytes())); nil != err {
		return err
	}

	buf, err = Package(stg.subthemeType, subthemeRoot)
	if nil != err {
		return err
	}
	if err := stg.creators["package"].Put(cfg.Theme.Id, bytes.NewReader(buf.Bytes())); nil != err {
		return err
	}

	return nil
}
