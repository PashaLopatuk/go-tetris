run:
	go run cmd/main.go

wasm:
	GOOS=js GOARCH=wasm go build -o ./bin/main.wasm ./cmd/main.go
	cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" ./bin
