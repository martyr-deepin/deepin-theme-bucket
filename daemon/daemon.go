package main

import (
	"os"
	"pkg.linuxdeepin.com/lib/dbus"
	dlog "pkg.linuxdeepin.com/lib/log"
)

var logger = dlog.NewLogger("deepin-download-service")

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
