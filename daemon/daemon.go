package main

import (
	"os"
	"pkg.deepin.io/lib/dbus"
	dlog "pkg.deepin.io/lib/log"
)

var logger = dlog.NewLogger("deepin-theme-bucket")

func main() {
	bs := NewBucketService()
	if err := bs.loadDBus(); nil != err {
		os.Exit(1)
	}
	dbus.DealWithUnhandledMessage()

	if err := dbus.Wait(); nil != err {
		logger.Error(err)
		os.Exit(1)
	}
	logger.Error("daemon Exit")
	os.Exit(0)
}
