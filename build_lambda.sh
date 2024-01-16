#!/bin/bash
# Build the lambda function and zip it up
#
# Usage: build_lambda.sh

mkdir -p build
rm -rf build/*

GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o build/bootstrap main.go

cd build || exit 1
SHORT_SHA=$(git rev-parse --short HEAD)
zip -r "lambda-$SHORT_SHA.zip" bootstrap

