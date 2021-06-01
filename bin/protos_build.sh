#!/bin/bash
protoc \
  -I ./ \
  --go_out ./ \
  --go-grpc_out ./ \
  --go_opt module=github.com/IronSavior/chat-grpc \
  --go-grpc_opt module=github.com/IronSavior/chat-grpc \
  ./protos/chat/*.proto

protoc \
  -I ./ \
  --go_out ./ \
  --go-grpc_out ./ \
  --go_opt module=github.com/IronSavior/chat-grpc \
  --go-grpc_opt module=github.com/IronSavior/chat-grpc \
  ./protos/auth/*.proto
