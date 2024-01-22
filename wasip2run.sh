# This requires TinyGo with WASI Preview 2 support
# https://github.com/dgryski/tinygo/tree/dgryski/wasi-preview-2

tinygo build -target=wasip2 -x -o main.wasm ./cmd/wasip2-test
wasm-tools component embed -w wasi:cli/command $(tinygo env TINYGOROOT)/lib/wasi-cli/wit/ main.wasm -o embedded.wasm
wasm-tools component new embedded.wasm -o component.wasm
wasmtime run --wasm component-model --env PWD --env USER --dir=. --dir=/tmp component.wasm ./LICENSE


# wasmtime serve --wasm component-model component.wasm


# old commands

# tinygo build -x -target=wasip2 -o main.wasm ./cmd/wasip2-test
# wasm-tools component embed -w wasi:cli/imports ~/src/bytecodealliance/wasmtime/crates/wasi/wit/ main.wasm -o embedded.wasm
# wasm-tools component new embedded.wasm -o component.wasm --adapt ./wasi_snapshot_preview1.command.wasm
# WASMTIME_LOG=trace WASMTIME_BACKTRACE_DETAILS=1 wasmtime run --wasm component-model --env KEY0=VALUE0 --env KEY1=VALUE1 --dir=. --dir=/tmp component.wasm arg1 arg2 arg3 arg4 arg5

# Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla suscipit commodo sapien vitae lobortis. Vivamus mattis odio eget velit commodo vehicula. Suspendisse velit odio, pulvinar id condimentum eget, scelerisque ut mi. Vivamus mi felis, mattis vitae orci nec, porta porta metus. Suspendisse potenti. Vestibulum sed cursus lectus. Sed eget justo vitae enim sodales dapibus. Phasellus vel quam vitae est iaculis maximus. Mauris sit amet gravida tellus. Etiam id commodo elit. Fusce rhoncus risus lectus, a maximus nisl facilisis id. Fusce ut efficitur felis. Aenean eu placerat lorem. Donec non tortor vel erat consectetur luctus.

# Mauris vestibulum nulla a sapien lobortis sagittis. Sed varius metus ac lorem laoreet, quis auctor augue dapibus. Morbi vehicula augue elit, gravida semper purus cursus in. Pellentesque facilisis magna rhoncus, laoreet urna id, placerat risus. Nulla euismod, sapien ut viverra ultricies, mi sapien dapibus mi, sit amet porta magna est nec dui. Nulla finibus nibh in arcu aliquet ullamcorper. Vestibulum lacinia dapibus urna id aliquam. Nulla sem urna, feugiat ac tellus quis, ullamcorper viverra nisi. Maecenas eu tempor ligula.

# Sed cursus at tellus sit amet euismod. Duis commodo enim convallis molestie euismod. Donec id metus leo. Fusce sem erat, lobortis eu hendrerit a, feugiat et est. Vivamus suscipit dolor hendrerit velit porta, et fringilla sem lobortis. Duis elementum ipsum sit amet erat congue auctor. Nulla feugiat nunc erat, quis rutrum augue mollis sit amet. Donec tellus risus, mollis et pellentesque eget, finibus a sem.

