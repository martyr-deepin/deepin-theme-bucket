.PHONY : test

CURDIR=$(shell pwd)
FIXGOPATH=$(CURDIR):$(GOPATH)
BIN_PATH=$(CURDIR)/bin

DDAEMON_SRC=$(CURDIR)/daemon

ifndef USE_GCCGO
    GOBUILD = go build
else
    LDFLAGS = $(shell pkg-config --libs gio-2.0)
    GOBUILD = go build -compiler gccgo -gccgoflags "${LDFLAGS}"
endif

build:
	cd $(DDAEMON_SRC)  && GOPATH=$(FIXGOPATH) ${GOBUILD} -o $(BIN_PATH)/deepin-theme-bucket-service

test:
	cd $(CURDIR)/bucket && GOPATH=$(FIXGOPATH) go test -v

install:
	@mkdir -p $(DESTDIR)/usr/share/dbus-1/services
	@mkdir -p $(DESTDIR)/etc/dbus-1/session.d
	cp -a $(CURDIR)/misc/dbus/com.deepin.theme.bucket.service.service $(DESTDIR)/usr/share/dbus-1/services/
	cp -a $(CURDIR)/misc/dbus/com.deepin.theme.bucket.service.conf $(DESTDIR)/etc/dbus-1/session.d/
	install -Dm755 $(BIN_PATH)/deepin-theme-bucket-service $(DESTDIR)/usr/lib/deepin-daemon/deepin-theme-bucket-service

clean:
	@-rm -rf bin/*
