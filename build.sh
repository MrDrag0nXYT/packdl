#!/bin/bash

PROJECT=packdl
VERSION=$1

go mod tidy

GOOS=windows
GOARCH=amd64

go build -ldflags "-s -w -X main.Version=%VERSION%" -o ${PROJECT}_${GOOS}_${GOARCH}.exe

GOOS=linux
GOARCH=amd64

go build -ldflags "-s -w -X main.Version=%VERSION%" -o ${PROJECT}_${GOOS}_${GOARCH}.bin
