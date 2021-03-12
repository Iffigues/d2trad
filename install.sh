#!/bin/sh

command -v go-bindata >/dev/null 2>&1 || {go get -u github.com/jteeuwen/go-bindata}
cd ./polaroid
go mod init polaroid
go get
cd ../staticserver
go mod init static
go-bindata -o myfile.go public/...
go get
cd ..
