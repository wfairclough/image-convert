#!/bin/bash
# Build the lambda function and zip it up
#
# Usage: build_lambda.sh

mkdir -p build
rm -rf build/*

GOOS=linux GOARCH=amd64 go build -o build/bootstrap main.go

cd build || exit 1
SHORT_SHA=$(git rev-parse --short HEAD)
zip -r "lambda-$SHORT_SHA.zip" bootstrap

