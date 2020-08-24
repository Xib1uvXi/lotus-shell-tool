#!/usr/bin/env bash

## $2 是日志

echo $1

lotus daemon >>$2 2>&1 &
