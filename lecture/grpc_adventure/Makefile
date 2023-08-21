.PHONY: protoc-deps
protoc-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# You need protoc commands by one of following:
#  apt install protobuf-compiler
#  apk add protoc
#  brew install protobuf
.PHONY: protoc
protoc: protoc-deps
	protoc --go_out=. --go-grpc_out=. welcome.proto

.PHONY: build
build:
	go build -o bin/server
