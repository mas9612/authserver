GOBIN=go
PROTOCBIN=protoc
BINNAME=authserver

.PHONY: all
all: dep test build-grpc build doc

.PHONY: build
build:
	CGO_ENABLED=0 $(GOBIN) build -o $(BINNAME) .

.PHONY: test
test:
	$(GOBIN) test -v ./...

.PHONY: build-grpc
build-grpc:
	$(PROTOCBIN) --go_out=plugins=grpc:. ./pkg/authserver/authserver.proto

.PHONY: doc
doc:
	$(PROTOCBIN) --doc_out=./doc --doc_opt=markdown,authserver.md ./pkg/authserver/authserver.proto

.PHONY: clean
clean:
	$(GOBIN) clean
	rm -f $(BINNAME)

.PHONY: dep
dep:
	dep ensure
