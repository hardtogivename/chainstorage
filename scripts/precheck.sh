#!/usr/bin/env bash

set -eo pipefail

# Check if protobuf is installed
if ! [ -x "$(command -v protoc)" ]; then
  echo 'Error: protobuf is not installed. Please refer https://grpc.io/docs/protoc-installation/ for installation instruction' >&2
  exit 1
fi

# Check if yq is installed
if ! [ -x "$(command -v yq)" ]; then
  echo 'Error: yq is not installed.' >&2
  exit 1
fi

# Check if protoc-gen-go-grpc is installed
if ! [ -x "$(command -v protoc-gen-go-grpc)" ]; then
  echo 'Error: protoc-gen-go-grpc is not installed.' >&2
  exit 1
fi

