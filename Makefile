.PHONY: default build release test clean

default: build

build: clean vet
	script/build

cross: clean vet
	script/build cross

docker: clean vet
	script/build docker

release: clean vet
	script/build docker
	script/release

run: build
	script/run

fmt:
	goimports -w src

vet:
	go vet ./src/...

test:
	script/test

clean:
	rm -rf bin
