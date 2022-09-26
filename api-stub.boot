#!/bin/bash
dir=$(dirname $0)
docker run -p 127.0.0.1:8081:80 --net=api-stub --name=api-stub -d mingsiro71/api-stub
docker run -p 6379 --net=api-stub --name=api-stub-redis -d redis
