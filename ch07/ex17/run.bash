#!/usr/bin/env bash

if [ ! -r "./sample.html" ]; then
  curl "https://www.w3.org/TR/2006/REC-xml11-20060816/" > ./sample.html
fi

cat sample.html | go run xmlselect.go  div.back div.div1 'a#dt-legacyenc'
#cat sample.html | go run xmlselect.go .body h2
