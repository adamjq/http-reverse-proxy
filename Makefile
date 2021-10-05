format:
	gofmt -s -w .

test:
	go test ./...

run-proxy:
	go run cmd/reverseproxy/main.go

run-a:
	go run cmd/service_a/main.go

run-b:
	go run cmd/service_b/main.go
