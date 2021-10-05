format:
	gofmt -s -w .

wire-install:
	go get github.com/google/wire/cmd/wire

generate: wire-install
	wire ./...
	go generate ./...

test:
	go test ./...

run-proxy:
	go run cmd/reverseproxy/main.go cmd/reverseproxy/wire_gen.go

run-a:
	go run cmd/service_a/main.go

run-b:
	go run cmd/service_b/main.go
