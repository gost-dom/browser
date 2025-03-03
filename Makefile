# Only one `gow` per terminal is allowed to use raw mode.
# Otherwise they conflict with each other.
RAW := $(if $(filter $(MAKELEVEL),0),-r,)

GOW_FLAGS := $(RAW)

# Expects an existing stable version of `gow`.
GOW := gow $(GOW_FLAGS)

.PHONY: main
main: codegen-watch

.PHONY: codegen-clean
codegen-clean:
	rm -f scripting/*_generated.go
	rm -f scripting/**/*_generated.go
	rm -f dom/*_generated.go
	rm -f html/*_generated.go

.PHONY: codegen-watch codegen-run
codegen-run:
	make -j2 codegen-watch codegen-build-watch

codegen-watch: codegen-clean
	$(GOW) -w ./internal/code-gen -S="Codegen done" -e="" generate ./...

.PHONY: codegen codegen-build codegen-build-watch

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
	# $(GOW) -c -e=go -e=js -e=html -v -w=./.. test -vet=off ./...
	gotestsum --format dots --watch ./... -- vet=off

.PHONY: test-dom
test-browser: 
	$(GOW) -s -w=./dom -w=./html -w=. test -vet=off . ./dom ./html

.PHONY: test-html
test-html: 
	cd html && ginkgo watch -vet=off

.PHONY: test-v8
test-v8: 
	# $(GOW) -s -e=go -e=js -e=html -w ./.. test -vet=off ./scripting/v8host
	gotestsum --format dots --watch --packages "./scripting/v8host ./internal/test/scripttests" -- -vet=off

test-scripting: 
	gotestsum --format dots --watch --packages "./scripting/... ./internal/test/scripttests" -- -vet=off

.PHONY: test-goja
test-goja:
	$(GOW) -c -e=go -e=js -e=html -w ./.. test -vet=off ./scripting/gojahost

 
.PHONY: ci ci-build
ci-build:
	go build -v ./...

ci: codegen ci-build test
	git diff --quiet HEAD
