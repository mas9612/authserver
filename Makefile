GOBIN=go
PROTOCBIN=protoc
AUTHSERVER=authserver
AUTHPROXY=authserver-proxy

.PHONY: all
all: dep test build-grpc build doc

.PHONY: build
build:
	CGO_ENABLED=0 $(GOBIN) build -o $(AUTHSERVER) ./cmd/authserver
	CGO_ENABLED=0 $(GOBIN) build -o $(AUTHPROXY) ./cmd/authserver-proxy

.PHONY: test
test:
	$(GOBIN) test -v ./...

.PHONY: build-grpc
build-grpc:
	$(PROTOCBIN) -I. -I$(GOPATH)/src -I/usr/local/include	\
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis	\
		--go_out=plugins=grpc:. ./pkg/authserver/authserver.proto
	$(PROTOCBIN) -I. -I$(GOPATH)/src -I/usr/local/include	\
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis	\
		--grpc-gateway_out=logtostderr=true:.	\
		./pkg/authserver/authserver.proto

.PHONY: doc
doc:
	$(PROTOCBIN) -I. -I$(GOPATH)/src -I/usr/local/include	\
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis	\
		--doc_out=./doc --doc_opt=markdown,authserver.md ./pkg/authserver/authserver.proto
	$(PROTOCBIN) -I. -I$(GOPATH)/src -I/usr/local/include	\
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis	\
		--swagger_out=logtostderr=true:. ./pkg/authserver/authserver.proto

.PHONY: clean
clean:
	$(GOBIN) clean
	rm -f $(BINNAME)

.PHONY: dep
dep:
	dep ensure
