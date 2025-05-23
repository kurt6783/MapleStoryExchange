#!/bin/bash

set -e
set -x

DATE=$(date +"%Y%m%d")
env="uat"
host="ec2-user@3.86.236.95"
pem_path="lineage.pem"
project_path="/Users/kurt.hsu/Desktop/MapleStoryExchange"
binary_name="main_$DATE"
local_bin_path="$project_path/bin/linux_amd64"
remote_path="/home/ec2-user/MapleStoryExchange"

mkdir -p "$local_bin_path"

docker run --rm -v "$project_path":/app -w /app \
    -e CGO_ENABLED=1 \
    -e GOOS=linux \
    -e GOARCH=amd64 \
    golang:1.23.8 \
    bash -c "apt-get update && apt-get install -y gcc && go mod tidy && go mod vendor && go build -ldflags='-s -w' -o ./bin/linux_amd64/$binary_name"

ssh -i "$pem_path" "$host" "kill -15 \$(ps -x | grep '$remote_path/main' | grep -v 'grep' | awk '{print \$1}')"

sleep 5

scp -i "$pem_path" "$local_bin_path/$binary_name" "$host":"$remote_path/main"

ssh -i "$pem_path" "$host" "nohup $remote_path/main >/dev/null 2>&1 &"
