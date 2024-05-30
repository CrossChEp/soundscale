#!/bin/bash

# Compile common protos itselfes
protos=$(find . -name '*.proto')

echo 'Compiling common protos...'
for proto in $protos; do
dir=$(dirname "$proto")
proto_name=$(basename "$proto")
protoc --go_out="$dir" --go-grpc_out="$dir" --proto_path="$dir" "$proto_name"
done

echo 'Compiling service specific protos...'
python3 generate_service_specific_protos.py