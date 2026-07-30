BREW_RAYLIB_PREFIX := $(shell brew --prefix raylib)

.PHONY: run build test

run:
	CGO_ENABLED=0 LD_LIBRARY_PATH="$(BREW_RAYLIB_PREFIX)/lib:$$LD_LIBRARY_PATH" DYLD_LIBRARY_PATH="$(BREW_RAYLIB_PREFIX)/lib:$$DYLD_LIBRARY_PATH" go run -tags "raylib raylib_no_embed" .

build:
	CGO_ENABLED=0 LD_LIBRARY_PATH="$(BREW_RAYLIB_PREFIX)/lib:$$LD_LIBRARY_PATH" DYLD_LIBRARY_PATH="$(BREW_RAYLIB_PREFIX)/lib:$$DYLD_LIBRARY_PATH" go build -tags "raylib raylib_no_embed" .

test:
	go test ./...
