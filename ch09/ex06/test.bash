#!/usr/bin/env bash

for i in `seq 0 5`; do
  env GOMAXPROCS=$(( 2 ** $i )) go test -v -bench=.
done
