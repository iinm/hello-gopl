#!/usr/bin/env bash

url="http://identicon.org/?t=test&s=1000"

go run fetchall.go $url
go run fetchall.go $url

cmp \?t=test\&s=1000_at_*
