#!/usr/bin/env bash

source .envrc

# 引数から
go run main.go 32 0.3048 0.453592

# 標準入力から
echo "32 0.3048 0.453592" | go run main.go
