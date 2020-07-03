#!/usr/bin/env sh

docker run --rm --privileged \
  -v $PWD:/go/src/khorevaa/r2gitsync \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -w /go/src/khorevaa/r2gitsync \
  goreleaser/goreleaser:latest release --snapshot --skip-publish --rm-dist