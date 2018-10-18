#!/usr/bin/env bash

#-> % free -m
#              total        used        free      shared  buff/cache   available
#Mem:           7855         340        7308          45         206        7248
#Swap:          8191        1218        6973

go run pipeline.go -length=2000000
# -> 一秒未満で完了

#go run pipeline.go -length=3000000
# -> メモリが足りず、swap発生して終わらない
