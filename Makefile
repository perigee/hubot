fmt:
	gofmt -s -w .

pb:
	protoc -I pb/ pb/service.protoc --go_out=plugins=grpc:pb
