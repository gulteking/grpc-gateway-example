
# GRPC-GATEWAY-EXAMPLE

Main goal of the repository is automatically creating a grpc and http server from an `example.proto` file and adding auth, logging and validation middlewares.


- `grpc-gateway` library is used for doing main job, it creates both grpc and http server(reverse proxy) .
- `grpc-middleware` for adding all middleware functionalities.
- `protoc-gen-validate` for message validations.
- `buf` user friendly proto cli


# Requirements
* Go 1.20: https://go.dev/doc/install
* protobuf compiler: https://grpc.io/docs/protoc-installation/
* buf cli: https://buf.build/docs/installation
* grpc-gateway: https://github.com/grpc-ecosystem/grpc-gateway
* protoc-gen-validate: https://github.com/bufbuild/protoc-gen-validate

# Compile&Run
* `buf mod update`
* `buf generate` in case of any change in the `example.proto` .
* `go mod vendor`
* `go run main.go`

# TODO

- [x] Basic example
- [x] Validation
- [ ] Auth middleware
- [ ] Logging middleware
- [ ] Good project structure
- [ ] Better validations(customized validation error messages)
- [ ] Few more functions to demonstrate get, delete and patch requests

