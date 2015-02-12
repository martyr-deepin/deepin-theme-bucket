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

func (bs *BucketService) loadDBus() error {
	if !lib.UniqueOnSession(BucketServiceDest) {
		return fmt.Errorf("There is aready a bucket service running")
	}

	if err := dbus.InstallOnSession(bs); nil != err {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func (bs *BucketService) GetMeta(themetype, id string) (string, error) {
	return bs.core.GetMeta(themetype, id)
}

func (bs *BucketService) GetThemePackage(uuid string) (string, error) {
	return bs.core.GetThemePackage(uuid)
}

func (bs *BucketService) GetWidgetPackage(suuid string) (string, error) {
	return bs.core.GetWidgetPackage(suuid)
}

func (bs *BucketService) GetIconPackage(suuid string) (string, error) {
	return bs.core.GetIconPackage(suuid)
}

func (bs *BucketService) GetCursorPackage(suuid string) (string, error) {
	return bs.core.GetCursorPackage(suuid)
}

func (bs *BucketService) GetWallpaper(duuid string) (string, error) {
	return bs.core.GetWallpaper(duuid)
}

func (bs *BucketService) GetFont(duuid string) (string, error) {
	return bs.core.GetFont(duuid)
}

func (bs *BucketService) GetMetaFile(themetype, id string) (string, error) {
	return bs.core.GetMetaFile(themetype, id)
}

func (bs *BucketService) GetThemePackageFile(uuid string) (string, error) {
	return bs.core.GetThemePackageFile(uuid)
}

func (bs *BucketService) GetWidgetPackageFile(suuid string) (string, error) {
	return bs.core.GetWidgetPackageFile(suuid)
}

func (bs *BucketService) GetIconPackageFile(suuid string) (string, error) {
	return bs.core.GetIconPackageFile(suuid)
}

func (bs *BucketService) GetCursorPackageFile(suuid string) (string, error) {
	return bs.core.GetCursorPackageFile(suuid)
}

func (bs *BucketService) GetWallpaperPackageFile(suuid string) (string, error) {
	return bs.core.GetWallpaperPackageFile(suuid)
}

func (bs *BucketService) GetFontPackageFile(suuid string) (string, error) {
	return bs.core.GetFontPackageFile(suuid)
}

func (bs *BucketService) GetWallpaperFile(duuid string) (string, error) {
	return bs.core.GetWallpaperFile(duuid)
}

func (bs *BucketService) GetFontFile(duuid string) (string, error) {
	return bs.core.GetFontFile(duuid)
}
