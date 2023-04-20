NAME := zeus-server
VERSION := 0.0.1
NOW     := $(shell date +%F_%T)
GO_VERSION := $(shell go version)


BINARY_NAME = zeus-server

GOOS := linux
GOARCHS := amd64

hello:
	echo "Hello"

release:
	sudo apt-get install zip -y
	export GO111MODULE=on
	mkdir -p bin
	$(foreach goos,$(GOOS),$(foreach goarch,$(GOARCHS),GOOS=$(goos) GOARCH=$(goarch) go build -ldflags "-w -s -X main.VERSION=$(VERSION) -X 'main.BUILD_TIME=$(NOW)' -X 'main.GO_VERSION=$(GO_VERSION)'" main.go;tar -czf bin/$(BINARY_NAME)-$(goos)-$(goarch).tar.gz main;))
	$(foreach goarch,$(GOARCHS),GOOS=windows GOARCH=$(goarch) go build  -ldflags "-w -s -X main.VERSION=$(VERSION) -X 'main.BUILD_TIME=$(NOW)' -X 'main.GO_VERSION=$(GO_VERSION)'" main.go ;mv main.exe zeus-server.exe;zip bin/$(BINARY_NAME)-windows-$(goarch).zip zeus-server.exe;)
	rm -rf main
	rm -rf zeus-server.exe

run:
	go run main.go

	
