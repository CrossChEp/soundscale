#!/usr/bin/env python
from array import array
import logging
import os
import subprocess
import sys
import yaml

CONFIG_FILE = "proto_dependencies.yml"


def validate_config_structure(conf: dict):
    for k in conf:
        assert conf[k].get("dependencies") is not None
        assert type(conf[k]["dependencies"]) == type(list())


print("Reading configs...")
# Load config
with open(CONFIG_FILE, "r") as f:
    try:
        config = yaml.safe_load(f)
    except yaml.YAMLError as exc:
        print(exc)
        sys.exit(1)

validate_config_structure(conf=config)

print("Compiling protos...")
# Compile protos
for service in config:
    print(f"[!] Compiling dependencies for {service}...")
    service_proto_path = (
        subprocess.run(
            f"find {service} -name 'proto*' -type d", shell=True, capture_output=True
        )
        .stdout.decode()
        .strip()
    )
    dependencies: list = config[service]["dependencies"]
    dependencies_protos = []
    for dep in dependencies:
        command = f"find {dep} -name *.proto"
        result = subprocess.run(command, shell=True, capture_output=True, check=True)
        proto = result.stdout.decode().strip()
        dependencies_protos += proto.split()
    for dep in dependencies_protos:
        dep_dir = os.path.dirname(dep)
        dep_name = os.path.basename(dep)
        command = f"protoc --go_out={service_proto_path} --go-grpc_out={service_proto_path} --proto_path={dep_dir} {dep_name}".strip()
        print(f"[x] {command}")
        result = subprocess.run(command, shell=True, capture_output=True)
        print(
            f"[ ] OK"
            if not result.stderr
            else f"[ERROR] {result.stderr.decode()}"
        )
        
print("Done!")