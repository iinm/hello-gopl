#!/usr/bin/env bash

# prepare url list
filename=top-1m.csv

if [ ! -f $filename ]; then
  echo 'Prepare url list...'
  curl -O http://s3.amazonaws.com/alexa-static/top-1m.csv.zip
  unzip top-1m.csv.zip
fi

urls=$(head -100 $filename | cut -d ',' -f 2 | awk '{ print "https://www."$1 }')
#echo $urls

go run ./fetchall.go $urls
