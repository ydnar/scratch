# This requires TinyGo with WASI Preview 2 support
# https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2

tinygo build -target=wasip2 -x -o main.wasm ./cmd/wasip2-test
wasm-tools component embed -w wasi:cli/command $(tinygo env TINYGOROOT)/lib/wasi-cli/wit/ main.wasm -o embedded.wasm
wasm-tools component new embedded.wasm -o component.wasm
wasmtime run --wasm component-model --env PWD --env USER --dir=. --dir=/tmp component.wasm ./LICENSE

# wasmtime serve --wasm component-model component.wasm
