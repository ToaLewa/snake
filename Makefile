RAYLIB_PREFIX ?= $(shell brew --prefix raylib)
export PKG_CONFIG_PATH := $(RAYLIB_PREFIX)/lib/pkgconfig:$(PKG_CONFIG_PATH)
export LD_LIBRARY_PATH := $(RAYLIB_PREFIX)/lib:$(LD_LIBRARY_PATH)
export CGO_LDFLAGS := -Wl,-rpath,$(RAYLIB_PREFIX)/lib $(CGO_LDFLAGS)

.PHONY: run build

run:
	go run .

build:
	go build .
