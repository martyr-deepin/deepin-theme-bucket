package bucket

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"pkg.linuxdeepin.com/lib/utils"
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
		if nil != err {
			return err
		}
		fmt.Println(hdr)
		switch hdr.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(dist+hdr.Name, 0644)
		case tar.TypeReg:
			f, _ := os.Create(dist + hdr.Name)
			io.Copy(f, r)
		}

	}
}
