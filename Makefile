all: build-wasm run

init-deps:
	@dep init -v

get-deps:
	@dep ensure -v

build-wasm:
	@GOOS=js GOARCH=wasm go build -o main.wasm

run:
	@goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
