.PHONY: run
run: build
	go run ./server

.PHONY: build
build:
	go mod tidy
	GOOS=js GOARCH=wasm go build -o  ./static/page.wasm ./src