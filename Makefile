# Only one `gow` per terminal is allowed to use raw mode.
# Otherwise they conflict with each other.
RAW := $(if $(filter $(MAKELEVEL),0),-r,)

GOW_FLAGS := $(RAW)

GO_TEST := GOEXPERIMENT=synctest go test

# Expects an existing stable version of `gow`.
GOW := go run github.com/mitranim/gow  $(GOW_FLAGS)

.PHONY: changes
changes:
	echo "Changes since \c"
	git tag --merged | sort -V | tail -1
	gorelease -base=`git tag --merged | sort -V | tail -1`

.PHONY: main
main: codegen-watch

.PHONY: codegen-clean
codegen-clean:
	rm -f **/*_generated.go
	rm -f scripting/internal/**/*_generated.go
	rm -f internal/interfaces/**/*_generated.go

.PHONY: codegen-watch codegen-run
codegen-run:
	nodemon  -w internal/code-gen -e go -x "make codegen || exit 1"

codegen-watch: codegen-clean
	$(GOW) -w ./internal/code-gen -S="Codegen done" -e="" generate ./...

.PHONY: codegen codegen-build codegen-build-watch

codegen-test:
	$(MAKE) -C internal/code-gen test

codegen-build:
	$(MAKE) -C internal/code-gen build

codegen-build-watch:
	$(MAKE) -C internal/code-gen build-watch

codegen: codegen-clean codegen-build
	go generate ./...

.PHONY: test test-watch test-browser test-v8 test-sobek test-wpt
test:
	$(GO_TEST) -v -vet=all ./...

test-wpt:
	$(MAKE) -C internal/test/wpt test

test-watch:
	gotestsum --format dots ./... -- vet=off || echo "Error"
	gotestsum --format dots --watch ./... -- vet=off

.PHONY: test-browser test-dom
test-dom:
	gotestsum --packages "`go list -deps ./dom | ./deps`" --format dots --watch

test-browser:
	gotestsum --packages "`go list -deps ./html | ./deps`" --format dots --watch

.PHONY: test-html
test-html:
	gotestsum --format dots --watch --packages "./scripting/html" -- -vet=off

.PHONY: test-v8
test-v8:
	gotestsum --format dots --watch --packages "./scripting/v8engine ./scripting/internal/scripttests" -- -vet=off

test-scripting:
	gotestsum --format dots --watch --packages "./scripting/... ./scripting/internal/scripttests" -- -vet=off

.PHONY: test-sobek
test-sobek:
	$(GOW) -c -e=go -e=js -e=html -w ./.. test -vet=off ./scripting/sobekengine

.PHONY: ci ci-build release
ci-build:
	go build -v ./...

wpt-watch:
	$(GOW) run ./internal/test/wpt

ci: codegen ci-build test test-wpt codegen-test
	git diff --quiet HEAD

release: ci
	pnpm run release
