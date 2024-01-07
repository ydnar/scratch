# This requires TinyGo with WASI Preview 2 support
# https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2

tinygo build -target=wasip2 -x -o main.wasm ./cmd/wasip2-test
wasm-tools component embed -w wasi:cli/command $(tinygo env TINYGOROOT)/lib/wasm/wasip2/wit/ main.wasm -o embedded.wasm
wasm-tools component new embedded.wasm -o component.wasm
wasmtime run --wasm component-model --env PWD --env USER --dir=. --dir=/tmp component.wasm arg1 arg2 arg3 arg4 arg5


# wasmtime serve --wasm component-model component.wasm


# old commands

# tinygo build -x -target=wasip2 -o main.wasm ./cmd/wasip2-test
# wasm-tools component embed -w wasi:cli/imports ~/src/bytecodealliance/wasmtime/crates/wasi/wit/ main.wasm -o embedded.wasm
# wasm-tools component new embedded.wasm -o component.wasm --adapt ./wasi_snapshot_preview1.command.wasm
# WASMTIME_LOG=trace WASMTIME_BACKTRACE_DETAILS=1 wasmtime run --wasm component-model --env KEY0=VALUE0 --env KEY1=VALUE1 --dir=. --dir=/tmp component.wasm arg1 arg2 arg3 arg4 arg5
