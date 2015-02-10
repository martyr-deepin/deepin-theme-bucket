package bucket

import (
	"io"
	"net/http"
)

type themeGenerator struct {
	client   *http.Client
	creators map[string]creator
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
