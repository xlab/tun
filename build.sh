#!/bin/sh

WORKSPACE=$GOPATH/src/github.com/xlab
PACKAGE=$WORKSPACE/tun
DOCK=$WORKSPACE/tun
NAME=tun

# compile the binary
goxc -arch amd64 -os linux -wd $PACKAGE -d $DOCK compile
# build the docker image
docker build -t xlab/$NAME $DOCK
# save docker image as tar archive
docker save -o $DOCK/$NAME.tar xlab/$NAME
# compress the archive
gzip -f $DOCK/$NAME.tar
