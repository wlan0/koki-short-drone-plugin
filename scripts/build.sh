#!/bin/bash 

set -e

src_dir=$(dirname $0)

#import version details
source $src_dir/version.sh

#build from parent dir
cd $src_dir/..

#create output dir if none exists
mkdir -p bin

#build a static go binary
CGO_ENABLED=0 go build -ldflags "-X main.GITCOMMIT=$VERSION -extldflags -static -w" -o bin/koki-short-drone-plugin

