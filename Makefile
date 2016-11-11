.PHONY: all
.DEFAULT_GOAL := help

GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

unexport http_proxy
unexport https_proxy
unexport all_proxy

help:
	$(info Available targets:)
	@grep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed 's/\:.*\#\#/\t\t/g'


fmt:	## Format all the go code
	gofmt -w $(GOFMT_FILES) 

pb:	## Generate go code from proto file
	protoc -I pb/ pb/service.protoc --go_out=plugins=grpc:pb

lanuch:	## Lanuch the whole local testing system
	docker-compose up
