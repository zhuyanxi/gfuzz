#!/bin/bash -e
cd "$(dirname "$0")"/..


GOMOD_DIR=$1
OUT_DIR=$2
shift 2

podman build -f docker/fuzzer/Dockerfile -t github.com/zhuyanxi/gfuzz:latest .

podman run --rm -it \
-v $GOMOD_DIR:/fuzz/target \
-v $OUT_DIR:/fuzz/output \
-v $(pwd)/tmp/pkgmod:/go/pkg/mod \
gfuzz:latest true /fuzz/target /fuzz/output $@
