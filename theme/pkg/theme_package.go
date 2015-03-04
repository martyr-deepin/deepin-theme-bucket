package themepkg

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"regexp"
)

type ThemePackage struct {
	Config      *tar.Reader
	MetaList    []*tar.Reader
	PackageList []*tar.Reader
}

func isConfig(name string) bool {
	validID := regexp.MustCompile(`^[^/][^/]*[/]theme.ini`)
	return validID.MatchString(name)
}

func isMeta(name string) bool {
	validID := regexp.MustCompile(`^[^/][^/]*[/]\.meta[/][^/]*`)
	return validID.MatchString(name)
}

func isMeta(name string) bool {
	validID := regexp.MustCompile(`^[^/][^/]*[/]\.meta[/][^/]*`)
	return validID.MatchString(name)
}

func NewThemePackage(r io.Reader) (*ThemePackage, error) {

func NewThemePackage(r io.Reader) (*ThemePackage, error) {
	tp := &ThemePackage{}
	gr, _ := gzip.NewReader(r)
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if nil != err {
			break
		}
		var buf []byte
		bufs := bytes.NewBuffer(buf)
		tw := tar.NewWriter(bufs)
		tw.WriteHeader(hdr)
		io.Copy(tw, tr)
		tw.Close()
		ftr := tar.NewReader(bufs)
		tp.PackageList = append(tp.PackageList, ftr)

		if isConfig(hdr.Name) {
			tp.Config = ftr
		}
	}
	return tp, nil
}
