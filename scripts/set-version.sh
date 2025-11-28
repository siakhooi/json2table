#!/bin/bash
set -e

# shellcheck disable=SC1091
. ./release.env

sed -i 'internal/version/version.go' -e 's@const json2tableVersion = ".*"@const json2tableVersion = "'"$RELEASE_VERSION"'"@g'
sed -i 'internal/version/version_test.go' -e 's@const expectedjson2tableVersion = ".*"@const expectedjson2tableVersion = "'"$RELEASE_VERSION"'"@g'
