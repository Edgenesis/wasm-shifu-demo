#!/bin/bash
cargo build --release --target wasm32-wasi
mv target/wasm32-wasi/release/js_func.wasm /reactor/src/js