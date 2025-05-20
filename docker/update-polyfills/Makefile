.PHONY: build-container rm-container
build-container: rm-container
	docker build . -t gost
	docker create --name gost-builder gost

rm-container:
	docker rm -f gost-builder
