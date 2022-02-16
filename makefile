APPLICATION      	  := new-go-application
LINUX            	  := build/${APPLICATION}-linux-amd64
DARWIN           	  := build/${APPLICATION}-darwin-amd64						
DOCKER_USER      	  ?= ""
DOCKER_PASS      	  ?= ""
VERSION			 	  ?= latest
GITHUB_SHA			  ?= 0.1.0

.PHONY: $(DARWIN)
$(DARWIN):
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ${DARWIN} *.go

.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${LINUX} *.go
	
.PHONY: release
release: test
	echo "${DOCKER_PASS}" | docker login -u "${DOCKER_USER}" --password-stdin
	docker build -t "${APPLICATION}" -f build/container/Dockerfile . --no-cache 
	docker tag  "${APPLICATION}:latest" "${DOCKER_USER}/${APPLICATION}:${VERSION}"
	docker tag  "${APPLICATION}:latest" "${DOCKER_USER}/${APPLICATION}:${GITHUB_SHA}"
	docker push "${DOCKER_USER}/${APPLICATION}:${VERSION}"
	docker push "${DOCKER_USER}/${APPLICATION}:${GITHUB_SHA}"

.PHONY: build
build: test
	go build -o ${APPLICATION}

.PHONY: test
tests:
	go test -cover -timeout 10s ./...
