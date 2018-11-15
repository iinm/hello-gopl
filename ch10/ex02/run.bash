#!/usr/bin/env bash

echo "# zip"
go run ls.go < test.zip

echo "# tar"
go run ls.go < test.tar
