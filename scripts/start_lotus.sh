#!/usr/bin/env bash

## $2 是日志

echo "$1"

lotus daemon >>"$1" 2>&1 &
