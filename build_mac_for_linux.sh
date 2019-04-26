#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64
go build ./automatic.go
mv ./automatic ./bin/automatic
git add .
git commit -m'-'
git push origin master

say "提交成功"