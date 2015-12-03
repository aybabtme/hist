#!/usr/bin/env bash

usage() {
    echo "USAGE: ./release.sh [version] [msg...]"
    exit 1
}

REVISION=$(git rev-parse HEAD)
GIT_TAG=$(git name-rev --tags --name-only $REVISION)
if [ "$GIT_TAG" = "" ]; then
    GIT_TAG="devel"
fi


VERSION=$1
if [ "$VERSION" = "" ]; then
    echo "Need to specify a version! Perhaps '$GIT_TAG'?"
    usage
fi

set -u -e

rm -rf /tmp/hist_build/

mkdir -p /tmp/hist_build/linux
GOOS=linux go build -ldflags "-X main.version=$VERSION" -o /tmp/hist_build/linux/hist ../
pushd /tmp/hist_build/linux/
tar cvzf /tmp/hist_build/hist_linux.tar.gz hist
popd

mkdir -p /tmp/hist_build/darwin
GOOS=darwin go build -ldflags "-X main.version=$VERSION" -o /tmp/hist_build/darwin/hist ../
pushd /tmp/hist_build/darwin/
tar cvzf /tmp/hist_build/hist_darwin.tar.gz hist
popd

go run ../main.go file < README.tmpl.md > ../README.md -var "version=$VERSION"
git add ../README.md
git commit -m 'release bump'

hub release create \
    -a /tmp/hist_build/hist_linux.tar.gz \
    -a /tmp/hist_build/hist_darwin.tar.gz \
    $VERSION

git push origin master
