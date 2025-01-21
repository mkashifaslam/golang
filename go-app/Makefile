clean:
	@[ -e web ] && echo "Removing web..." && rm -rf web && echo "web removed" || echo "web does not exist"
	@[ -e hello ] && echo "Removing hello..." && rm hello && echo "hello removed" || echo "hello does not exist"
build: clean
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o hello

run: build
	./hello
