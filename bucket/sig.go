package bucket

import (
	"github.com/Iceyer/go-sdk/upyun/form"
	"io/ioutil"
	"net/http"
)

func getBody(url string) string {
	rsp, err := http.Get(url)
	if nil != err {
		return ""
	}
	data, _ := ioutil.ReadAll(rsp.Body)
	return string(data)

}

type remoteSignature struct {
}

func (rs *remoteSignature) SigFile(p form.Policy) string {
	api := "https://api.linuxdeepin.com/crop/upyun/sig/sigle"
	url := api + "?" + p.UrlEncode()
	return getBody(url)
}

func (rs *remoteSignature) SigBolocks(p form.Policy) string {
	api := "https://api.linuxdeepin.com/crop/upyun/sig/blocks"
	url := api + "?" + p.UrlEncode()
	return getBody(url)
}
