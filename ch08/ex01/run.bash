#!/usr/bin/env bash

trap "kill 0" EXIT

TZ=US/Eastern go run clock2.go -port 8010 &
TZ=Asia/Tokyo go run clock2.go -port 8020 &
TZ=Europe/London go run clock2.go -port 8030 &

go run clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
