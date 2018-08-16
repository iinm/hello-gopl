#!/usr/bin/env bash

for i in $(seq 1 300); do
    url="https://xkcd.com/$i/info.0.json"
    echo $url
    curl --silent -o data/$i.json $url
    sleep 0.5
done
