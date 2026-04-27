#!/usr/bin/env bash

set -e

wget https://raw.githubusercontent.com/saichler/l8types/refs/heads/main/proto/api.proto
wget https://raw.githubusercontent.com/saichler/l8common/refs/heads/main/proto/l8common.proto

docker run --user "$(id -u):$(id -g)" -e PROTO=spring-categories.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=spring-listings.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=spring-bids.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=spring-deals.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest
docker run --user "$(id -u):$(id -g)" -e PROTO=spring-reviews.proto --mount type=bind,source="$PWD",target=/home/proto/ -i saichler/protoc:latest

rm api.proto l8common.proto

rm -rf ../go/types
mkdir -p ../go/types
rm -rf ./types/l8common
mv ./types/* ../go/types/.
rm -rf ./types

rm -rf *.rs

cd ../go
find . -name "*.go" -type f -exec sed -i 's|"./types/l8api"|"github.com/saichler/l8types/go/types/l8api"|g' {} +
find . -name "*.go" -type f -exec sed -i 's|"./types/l8common"|"github.com/saichler/l8common/go/types/l8common"|g' {} +
find . -name "*.go" -type f -exec sed -i 's|"./types/spring"|"github.com/saichler/l8spring/go/types/spring"|g' {} +
