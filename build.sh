#!/bin/sh

EXEC=tun
PACKAGE=github.com/xlab/tun
WORKDIR=$GOPATH/src/$PACKAGE
DOCKOUT=$WORKDIR/docker

# compile the binary (uncomment the option)
# goxc -q -arch amd64 -os linux -wd $WORKDIR -o="$DOCKOUT/bin/$EXEC" compile
# gobldock -o $DOCKOUT/bin/$EXEC $PACKAGE
exit 0

# build the docker image
cp $WORKDIR/Dockerfile.buildsh $DOCKOUT/Dockerfile
docker build -t xlab/$EXEC $DOCKOUT
# save docker image as tar archive
docker save -o $DOCKOUT/$EXEC.tar xlab/$EXEC
# compress the archive
gzip -f $DOCKOUT/$EXEC.tar
