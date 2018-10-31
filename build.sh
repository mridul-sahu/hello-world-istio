#!/bin/sh
protoc -I ./protos --go_out=plugins=grpc:./protos ./protos/*.proto
go get -v
