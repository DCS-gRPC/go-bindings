# DCS-gRPC Go language bindings

[![Go Reference](https://pkg.go.dev/badge/github.com/DCS-gRPC/go-bindings.svg)](https://pkg.go.dev/github.com/DCS-gRPC/go-bindings)

## Introduction

This project provides auto-generated Go sources for use alongside the
[DCS-gRPC](https://github.com/DCS-gRPC/rust-server) project.

## Examples

Please see the [examples/](examples/) directory for some example usage of the Go
bindings.

## Development

### Build Dependencies

- Go `>=1.17`
- protoc [`>=3.20`](https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.1)
- protoc-gen-go `1.28.0`
- protoc-gen-go-grpc `1.2.0`

`protoc-gen-go` and `protoc-gen-go-grpc` versions are tracked via `go.mod` to ensure consistent client generation and can be installed using `make tools`.

### Generating client

- Update submodule ref to desired version of [`rust-server`](https://github.com/DCS-gRPC/rust-server) to generate client from
- `make generate`
