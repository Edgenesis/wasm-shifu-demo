# Image for doing the build work

FROM rust:latest

RUN rustup target add wasm32-wasi

RUN mkdir /reactor
ADD . /reactor

RUN cd /reactor && \
cargo build --release --target wasm32-wasi && \
rm target/wasm32-wasi/release/js_func.*

WORKDIR /reactor
