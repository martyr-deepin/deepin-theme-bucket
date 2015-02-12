package bucket

import (
	"io"
	"net/http"
)

type themeGenerator struct {
	client   *http.Client
	creators map[string]creator
}

func (g *themeGenerator) GetURL(datatype string, id string) (string, error) {
	return g.creators[datatype].GetURL(id)
}

func (g *themeGenerator) Get(datatype string, id string) (io.ReadCloser, error) {
	return g.creators[datatype].Get(id)
}

func (wpg *themeGenerator) Put(datatype string, r io.Reader) error {
	return nil
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

func newIconGenrator() generator {
	g := &themeGenerator{
		client:   &http.Client{},
		creators: map[string]creator{},
	}

	tpls := map[string]string{
		"meta":      "/subtheme/icon/%s/meta.tar.gz",
		"config":    "/subtheme/icon/%s/theme.ini",
		"thumbnail": "/subtheme/icon/%s/thumbnail.png",
		"package":   "/subtheme/icon/%s/package.tar.gz",
	}

	for k, v := range tpls {
		g.creators[k] = &urlCreator{
			client:      g.client,
			urlTemplate: v,
		}
	}

	return g
}

func newWidgetGenrator() generator {
	g := &themeGenerator{
		client:   &http.Client{},
		creators: map[string]creator{},
	}

	tpls := map[string]string{
		"meta":      "/subtheme/widget/%s/meta.tar.gz",
		"config":    "/subtheme/widget/%s/theme.ini",
		"thumbnail": "/subtheme/widget/%s/thumbnail.png",
		"package":   "/subtheme/widget/%s/package.tar.gz",
	}

	for k, v := range tpls {
		g.creators[k] = &urlCreator{
			client:      g.client,
			urlTemplate: v,
		}

	}

	return g
}

func newCursorGenrator() generator {
	g := &themeGenerator{
		client:   &http.Client{},
		creators: map[string]creator{},
	}

	tpls := map[string]string{
		"meta":      "/subtheme/cursor/%s/meta.tar.gz",
		"config":    "/subtheme/cursor/%s/theme.ini",
		"thumbnail": "/subtheme/cursor/%s/thumbnail.png",
		"package":   "/subtheme/cursor/%s/package.tar.gz",
	}

	for k, v := range tpls {
		g.creators[k] = &urlCreator{
			client:      g.client,
			urlTemplate: v,
		}
	}

	return g
}
