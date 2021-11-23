#!/usr/bin/env bash

VERSION=2.0

docker build . -t hellodapr-go:${VERSION}

docker tag hellodapr-go:${VERSION} nickdala/hellodapr-go:${VERSION}

docker push nickdala/hellodapr-go:${VERSION}