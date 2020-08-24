#!/usr/bin/env bash

## $2 是日志
echo "$1"
nohup lotus-miner run >>"$2" 2>&1 &
