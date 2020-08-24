#!/usr/bin/env bash

mkdir -p $7
export PATH=$PATH:~/tools/filecoin/calibration
nohup lotus-worker --worker-repo=$1 run --listen=$2 --addpiece=$3 --precommit1=$4 --precommit2=$5 --commit=$6 >>$7 2>&1 &
