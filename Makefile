all: gen test 

test:
	@echo "Testing..."
	@go test ./...

gen:
	@go generate ./...

clean:
	@rm -rf bin

deps: dep_protoc

dep_protoc:
	@go get -u github.com/golang/protobuf/protoc-gen-go
