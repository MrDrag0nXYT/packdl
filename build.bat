@echo off

SET PROJECT=packdl
SET VERSION=%1

go mod tidy

set GOOS=windows
SET GOARCH=amd64

go build -ldflags "-s -w -X main.Version=%VERSION%" -o %PROJECT%_%GOOS%_%GOARCH%.exe

set GOOS=linux
SET GOARCH=amd64

go build -ldflags "-s -w -X main.Version=%VERSION%" -o %PROJECT%_%GOOS%_%GOARCH%.bin
