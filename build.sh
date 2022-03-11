#!/bin/sh
GOOS=js GOARCH=wasm go build -ldflags '-s -w' -o static/main.wasm ./main