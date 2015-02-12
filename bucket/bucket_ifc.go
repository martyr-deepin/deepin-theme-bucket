package bucket

func (b *Bucket) GetMeta(themetype, id string) (string, error) {
	return b.getURL(themetype, "meta", id)
}

func (b *Bucket) GetThemePackage(uuid string) (string, error) {
	return b.getURL("theme", "package", uuid)
}

func (b *Bucket) GetWidgetPackage(suuid string) (string, error) {
	return b.getURL("widget", "package", suuid)
}

func (b *Bucket) GetIconPackage(suuid string) (string, error) {
	return b.getURL("icon", "package", suuid)
}

func (b *Bucket) GetCursorPackage(suuid string) (string, error) {
	return b.getURL("cursor", "package", suuid)
}

func (b *Bucket) GetWallpaper(duuid string) (string, error) {
	return b.getURL("wallpaper", "data", duuid)
}

func (b *Bucket) GetFont(duuid string) (string, error) {
	return b.getURL("font", "data", duuid)
}

func (b *Bucket) GetMetaFile(themetype, id string) (string, error) {
	return b.getFile(themetype, "meta", id)
}

func (b *Bucket) GetThemePackageFile(uuid string) (string, error) {
	return b.getFile("theme", "package", uuid)
}

func (b *Bucket) GetWidgetPackageFile(suuid string) (string, error) {
	return b.getFile("widget", "package", suuid)
}

func (b *Bucket) GetIconPackageFile(suuid string) (string, error) {
	return b.getFile("icon", "package", suuid)
}

func (b *Bucket) GetCursorPackageFile(suuid string) (string, error) {
	return b.getFile("cursor", "package", suuid)
}

func (b *Bucket) GetWallpaperPackageFile(suuid string) (string, error) {
	return b.getFile("wallpaper", "package", suuid)
}

func (b *Bucket) GetFontPackageFile(suuid string) (string, error) {
	return b.getFile("font", "package", suuid)
}

func (b *Bucket) GetWallpaperFile(duuid string) (string, error) {
	return b.getFile("wallpaper", "data", duuid)
}

func (b *Bucket) GetFontFile(duuid string) (string, error) {
	return b.getFile("font", "data", duuid)
}
