#!/usr/bin/env bash
set -euxo pipefail

echo "--- go get for ${GO_VERSION}"
go get -v -t ./...

echo "--- go test for ${GO_VERSION}"
set +e
export OUT_FILE="build/test-report.out"
mkdir -p build
go test -v -race ./... 2>&1 | tee ${OUT_FILE}
status=$?

go get -v -u github.com/jstemmer/go-junit-report
go-junit-report > "build/junit-${GO_VERSION}.xml" < ${OUT_FILE}

exit ${status}
