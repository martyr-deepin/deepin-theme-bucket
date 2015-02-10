package bucket

import (
	"fmt"
	"io"
	"net/http"
)

var (
	BucketHost = "http://theme-store.b0.upaiyun.com/"
)

type generator interface {
	Get(datatype string, id string) (io.ReadCloser, error)
}

type creator interface {
	Get(id string) (io.ReadCloser, error)
}

type urlCreator struct {
	client      *http.Client
	urlTemplate string
}

func (uc *urlCreator) Get(id string) (io.ReadCloser, error) {
	url := BucketHost + fmt.Sprintf(uc.urlTemplate, id)
	rsp, err := uc.client.Get(url)
	if nil != err {
		return nil, err
	}
	if http.StatusOK != rsp.StatusCode {
		return nil, fmt.Errorf(rsp.Status + url)
	}
	return rsp.Body, nil
}
