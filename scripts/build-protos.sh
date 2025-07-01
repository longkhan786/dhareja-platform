#!/bin/bash

PROTO_DIR=./proto
OUT_DIR=./services/AgroService/gen

mkdir -p $OUT_DIR

protoc \
  --proto_path=$PROTO_DIR \
  --go_out=$OUT_DIR \
  --go-grpc_out=$OUT_DIR \
  $PROTO_DIR/agroTraders.proto