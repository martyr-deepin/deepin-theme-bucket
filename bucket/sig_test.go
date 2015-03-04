package bucket

import (
	"crypto/rand"
	"github.com/Iceyer/go-sdk/upyun/form"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

var (
	BucKet = "theme-store"
)

func TestForm(t *testing.T) {
	Convey("Test Form", t, func() {
		Convey("Test Post Sinle File", func() {
			uf := form.NewUpForm(BucKet, &remoteSignature{})
			err := uf.PostFile("sig.go", "/img.png")
			So(err, ShouldBeNil)
		})

		Convey("Test MutiPart Post Sinle File", func() {
			filepath, err := createTestFile(4*1024*1024 + 789)
			So(err, ShouldBeNil)
			uf := form.NewUpForm("theme-store", &remoteSignature{})
			err = uf.SlicePostFile(filepath, "/big.zip")
			So(err, ShouldBeNil)
		})
	})
}

func createTestFile(size int64) (string, error) {
	tmpfilepath := os.TempDir() + "/.upyun-tmp-test-createTest-File"
	file, err := os.Create(tmpfilepath)
	if nil != err {
		return "", err
	}

	buf := make([]byte, size)
	_, err = rand.Read(buf)
	if nil != err {
		return "", err
	}
	file.Write(buf)
	file.Close()
	return tmpfilepath, nil
}
