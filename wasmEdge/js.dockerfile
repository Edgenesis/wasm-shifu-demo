# Build the manager binary
FROM rust:latest as builder

RUN rustup target add wasm32-wasi

WORKDIR /reactor
COPY js-func /reactor

RUN cargo build --release --target wasm32-wasi && \
    mv target/wasm32-wasi/release/js_func.wasm js_func.wasm

FROM edgehub/wasm-demo:v0.0.1

WORKDIR /worker

COPY --from=builder /reactor/js_func.wasm js_func.wasm

