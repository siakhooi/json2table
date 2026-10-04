# List available recipes
default:
    @just --list

# Remove build outputs
clean:
    rm -rf bin dist

# Run golangci-lint
golangci-lint:
    golangci-lint run

# Run tests and write coverage reports
test:
    scripts/test.sh

# Lint and test
ci:
    scripts/ci.sh

# Cross-compile into bin/ and build snapshot archives, debs, and rpms into dist/
build:
    scripts/build.sh
    scripts/goreleaser.sh snapshot

# Clean, lint, test, and build
all: clean ci build

# Create a GitHub release from release.env
release:
    scripts/create-release.sh

# Build snapshot archives, debs, and rpms in dist/ without uploading
go-release:
    scripts/goreleaser.sh snapshot

# Watch the current GitHub Actions run
commit-watch:
    gh run watch

# Create a GitHub release and watch the Actions run
release-watch: release
    gh run watch

json2table := "bin/json2table-linux-amd64"

# Show CLI help
run-help:
    {{ json2table }} -h

# Print the version
run-version:
    {{ json2table }} -v

# Print build info
run-build:
    {{ json2table }} --build

# Run json2table with the given arguments
run *args:
    {{ json2table }} {{ args }}

# Exercise the common CLI paths
run-common:
    {{ json2table }} -s ./samples/spec1.json ./samples/data1.json
    {{ json2table }} -s ./samples/spec2.json ./samples/data1.json
    {{ json2table }} -c 'id,desc,url,display.name' ./samples/data2.json
    {{ json2table }} --spec ./samples/spec1.json ./samples/data1.json
    cat ./samples/data1.json | {{ json2table }} -s ./samples/spec1.json
    JSON2TABLE_SPEC_FILE=./samples/spec1.json {{ json2table }} ./samples/data1.json
    cat ./samples/data1.json | JSON2TABLE_SPEC_FILE=./samples/spec1.json {{ json2table }}
    JSON2TABLE_SPEC='{"dataPath":"$.data2","columns":[{"path":"id","title":"ID"},{"path":"display.name"}]}' {{ json2table }} ./samples/data1.json

# Exercise invalid CLI paths
run-invalid:
    -{{ json2table }}
    -{{ json2table }} -s ./samples/spec1.json
    -{{ json2table }} ./samples/data1.json ./samples/data1.json
    -{{ json2table }} -s ./samples/spec.json ./samples/data1.json ./samples/data1.json

# Run with no arguments
run-no-arguments-1:
    -{{ json2table }}

# Run with a spec and no data file
run-no-arguments-2:
    -{{ json2table }} -s ./samples/spec1.json

# Run with too many data files
run-too-many-arguments-1:
    -{{ json2table }} ./samples/data1.json ./samples/data1.json

# Run with a spec and too many data files
run-too-many-arguments-2:
    -{{ json2table }} -s ./samples/spec.json ./samples/data1.json ./samples/data1.json

# Print samples/data1.json with spec1
run-1:
    {{ json2table }} -s ./samples/spec1.json ./samples/data1.json

# Print samples/data1.json with spec2
run-1a:
    {{ json2table }} -s ./samples/spec2.json ./samples/data1.json

# Print selected columns from samples/data2.json
run-1c:
    {{ json2table }} -c 'id,desc,url,display.name' ./samples/data2.json

# Print samples/data1.json with the long spec flag
run-2:
    {{ json2table }} --spec ./samples/spec1.json ./samples/data1.json

# Print samples/data1.json from stdin
run-3:
    cat ./samples/data1.json | {{ json2table }} -s ./samples/spec1.json

# Print samples/data1.json using JSON2TABLE_SPEC_FILE
run-4:
    JSON2TABLE_SPEC_FILE=./samples/spec1.json {{ json2table }} ./samples/data1.json

# Print stdin using JSON2TABLE_SPEC_FILE
run-5:
    cat ./samples/data1.json | JSON2TABLE_SPEC_FILE=./samples/spec1.json {{ json2table }}

# Print samples/data1.json using JSON2TABLE_SPEC
run-6:
    JSON2TABLE_SPEC='{"dataPath":"$.data2","columns":[{"path":"id","title":"ID"},{"path":"display.name"}]}' {{ json2table }} ./samples/data1.json
