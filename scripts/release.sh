#!/bin/sh

set -e

src_dir=$(dirname $0)

cd $src_dir/..

SHORT_VERSION=v0.3.0

GOOS=linux GOARCH=amd64 ./scripts/build.sh

docker build --build-arg SHORT_VERSION=${SHORT_VERSION} -t kokster/kubeci-plugin:${SHORT_VERSION} .
