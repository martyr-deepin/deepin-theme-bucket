package main

import (
	"../bucket"
	"fmt"
	"pkg.linuxdeepin.com/lib"
	"pkg.linuxdeepin.com/lib/dbus"
)

const (
	BucketServiceDest = "com.deepin.theme.bucket.service"
	BucketServicePath = "/com/deepin/theme/bucket/service"
	BucketServiceIfc  = "com.deepin.theme.bucket.service"
)

type BucketService struct {
	core *bucket.Bucket
}

func NewBucketService() *BucketService {
	return &BucketService{
		core: bucket.NewBucket(),
	}
}

func (bs *BucketService) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		Dest:       BucketServiceDest,
		ObjectPath: BucketServicePath,
		Interface:  BucketServiceIfc,
	}
}

func (bs *BucketService) Get(themeType string, datatype string, id string, filepath string) error {
	return bs.core.GetFile(themeType, datatype, id, filepath)
}

func (bs *BucketService) loadDBus() error {
	if !lib.UniqueOnSession(BucketServiceDest) {
		return fmt.Errorf("There is aready a bucket service running")
	}

	if err := dbus.InstallOnSession(bs); nil != err {
		return fmt.Errorf("%v", err)
	}
	return nil
}
