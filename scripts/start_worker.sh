#!/usr/bin/env bash

## $2 是日志
echo "$1"
nohup lotus-worker --worker-repo=$2 run --listen=$3 --addpiece=$4 --precommit1=$5 --precommit2=$6 --commit=$7 >>$8 2>&1 &
