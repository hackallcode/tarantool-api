#!/usr/bin/env bash

docker run --rm -d \
    -e TARANTOOL_USER_NAME=kv_storage \
    -e TARANTOOL_USER_PASSWORD=kv_storage \
    -p 3301:3301 \
    -v `pwd`/app.lua:/opt/tarantool/app.lua \
    --name tarantool \
    tarantool/tarantool:2.2.1 \
    tarantool /opt/tarantool/app.lua
