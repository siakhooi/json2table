#!/bin/bash
set -e

# shellcheck disable=SC1091
. ./release.env

sed -i 'internal/versioninfo/versioninfo.go' -e 's@const json2tableVersion = ".*"@const json2tableVersion = "'"$RELEASE_VERSION"'"@g'
sed -i 'internal/versioninfo/versioninfo_test.go' -e 's@const expectedjson2tableVersion = ".*"@const expectedjson2tableVersion = "'"$RELEASE_VERSION"'"@g'
