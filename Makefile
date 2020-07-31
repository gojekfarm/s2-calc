all: build-wasm run

get-deps:
	@go mod tidy -v

build-wasm:
	@GOOS=js GOARCH=wasm go build -o main.wasm

run:
	@goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
