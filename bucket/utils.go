package bucket

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"pkg.linuxdeepin.com/lib/utils"
	"strings"
)

var _randDir = os.TempDir() + "/" + utils.GenUuid()

func init() {
	os.Mkdir(_randDir, 0755)
}

func randTmpPath() string {
	return _randDir + "/" + utils.GenUuid()
}

func Extrat(r *tar.Reader, dist string) error {
	for {
		hdr, err := r.Next()
		if io.EOF == err {
			return nil
		}
		if nil != err {
			return err
		}
		switch hdr.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(dist+"/"+hdr.Name, 0755)
		case tar.TypeReg:
			fullname := dist + "/" + hdr.Name
			os.MkdirAll(filepath.Dir(fullname), 0755)
			f, _ := os.Create(fullname)
			io.Copy(f, r)
			f.Close()
		}

	}
}

func Package(prefix, root string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	gw := gzip.NewWriter(buf)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	if err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		hdr := &tar.Header{
			Name: prefix + strings.Replace(path, root, "", 1),
			Size: f.Size(),
			Mode: 0644,
		}
		tw.WriteHeader(hdr)
		data, err := os.Open(path)
		if nil != err {
			return err
		}
		io.Copy(tw, data)
		return nil
	}); nil != err {
		return nil, err
	}
	tw.Close()
	gw.Close()
	return buf, nil
}
