APPLICATION      	  := new-go-application
LINUX            	  := bin/${APPLICATION}-linux-amd64
DARWIN           	  := bin/${APPLICATION}-darwin-amd64						
DOCKER_USER      	  ?= ""
DOCKER_PASS      	  ?= ""
VERSION			 	  ?= latest
GITHUB_SHA			  ?= 0.1.0
BINARY_FOLDER		  ?= bin

.PHONY: help ## Display this help screen
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | cut -d: -f2

.PHONY: darwin ## Build the binary for mac
darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ${DARWIN} *.go

.PHONY: linux ## Build the binary for linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${LINUX} *.go
	
.PHONY: release ## Tests, package the image and push it to docker registry
release: test
	echo "${DOCKER_PASS}" | docker login -u "${DOCKER_USER}" --password-stdin
	docker build -t "${APPLICATION}" -f build/container/Dockerfile . --no-cache 
	docker tag  "${APPLICATION}:latest" "${DOCKER_USER}/${APPLICATION}:${VERSION}"
	docker tag  "${APPLICATION}:latest" "${DOCKER_USER}/${APPLICATION}:${GITHUB_SHA}"
	docker push "${DOCKER_USER}/${APPLICATION}:${VERSION}"
	docker push "${DOCKER_USER}/${APPLICATION}:${GITHUB_SHA}"

.PHONY: test ## Launch the unit test
tests:
	go test -cover -timeout 10s ./...

.PHONY: clean ## Remove previous build
clean: 
	go clean .
	rm -f ${BINARY_FOLDER}/*
