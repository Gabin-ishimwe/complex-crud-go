build:
	@go build -o bin/complex-crud cmd/main.go
test:
	@go test -v ./...
run: build
	@./bin/complex-crud