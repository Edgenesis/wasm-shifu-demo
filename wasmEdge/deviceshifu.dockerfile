FROM rust:latest as wasmbuilder

RUN rustup target add wasm32-wasi

WORKDIR /reactor
COPY wasmEdge/js-func /reactor

RUN cargo build --release --target wasm32-wasi && \
    mv target/wasm32-wasi/release/js_func.wasm js_func.wasm

FROM golang:1.18 as gobuilder

WORKDIR /worker

ENV GOPROXY=https://goproxy.cn,direct

COPY wasmEdge/main.go main.go
COPY wasmEdge/go.mod go.mod
COPY wasmEdge/go.sum go.sum
COPY wasmEdge/install.sh install.sh

RUN go mod download

RUN bash install.sh -v 0.10.0
RUN /bin/bash -c "source ~/.wasmedge/env && go build"

COPY --from=wasmbuilder /reactor/js_func.wasm js_func.wasm

EXPOSE 8080/tcp

ENTRYPOINT ["./main"]
