#!/bin/sh

set -e

rm -rf api/gen
mkdir -p api/gen

for dir in `find api/proto/api -type d`; do
	count=`find $dir -maxdepth 1 -name *.proto| wc -l`
	# protoファイルが存在するディレクトリのみ実行
	if [ $count -ne "0" ]; then
		protoc -I=api/proto/api \
				-I=$(go env GOPATH)/pkg/mod/github.com/gogo/protobuf@v1.3.2 \
				-I=$(go env GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 \
				--gofast_out=api \
				--validate_out="lang=go:api" \
				$dir/*.proto
		dirname=${dir##*/}
	fi
done

