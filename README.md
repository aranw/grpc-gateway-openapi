gRPC Gateway and OpenAPI Demo
=============================

This repository contains a working demo of using gRPC, gRPC Gateway and protoc plugins to generated OpenAPI V3 JSON/YAML files for generating API documentation. 

To generate the gRPC code and API documentation you need to run the following commands:

    1. Generate Protocol Buffers `buf generate --template buf.gen.yaml`
    2. Generate OpenAPI documentation `./scripts/build-docs.sh` 
    3. Convert OpenAPI V2 to OpenAPI V3 `go run ./v3gen`
    4. Generate redoc static documentation `redoc-cli build api/docs/v3/openapi3.json`

