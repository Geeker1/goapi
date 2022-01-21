#!/bin/bash

set -e
set -x

go build main.go

docker build -t goapi:1.0 .
