#!/bin/bash
port=$1
if [ -z "$port" ];then
    echo "port is required." 1>&2
    exit 1
fi
dir=$(dirname $(readlink -f $0))
docker run -p 127.0.0.1:${port}:80 --net=api-stub --name=api-stub --env-file ${dir}/.env -d mingsiro71/api-stub
docker run -p 6379 --net=api-stub --name=api-stub-redis -d redis
