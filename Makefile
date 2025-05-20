# Only one `gow` per terminal is allowed to use raw mode.
# Otherwise they conflict with each other.
RAW := $(if $(filter $(MAKELEVEL),0),-r,)

GOW_FLAGS := $(RAW)

# Expects an existing stable version of `gow`.
GOW := gow $(GOW_FLAGS)

.PHONY: changes
changes:
	echo "Changes since \c"
	git tag --merged | sort -V | tail -1
	gorelease -base=`git tag --merged | sort -V | tail -1`

.PHONY: main
main: codegen-watch

.PHONY: codegen-clean
codegen-clean:
	rm -f scripting/*_generated.go
	rm -f scripting/**/*_generated.go
	rm -f dom/*_generated.go
	rm -f internal/**/*_generated.go
	rm -f html/*_generated.go

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

.PHONY: test test-watch test-browser test-v8 test-goja
test: 
	go test -v -vet=all ./...

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
	cd html && ginkgo watch -vet=off

.PHONY: test-v8
test-v8: 
	gotestsum --format dots --watch --packages "./scripting/v8host ./internal/test/scripttests" -- -vet=off

test-scripting: 
	gotestsum --format dots --watch --packages "./scripting/... ./internal/test/scripttests" -- -vet=off

.PHONY: test-goja
test-goja:
	$(GOW) -c -e=go -e=js -e=html -w ./.. test -vet=off ./scripting/gojahost
 
.PHONY: ci ci-build release
ci-build:
	go build -v ./...

ci: codegen ci-build test codegen-test
	git diff --quiet HEAD

release: ci
	pnpm run release
